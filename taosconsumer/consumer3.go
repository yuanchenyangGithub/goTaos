package taosconsumer

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/Shopify/sarama"

	confs "disutaos/config"
	tbatter "disutaos/modules"
)

var conf = confs.GetConfig()

type consumerGroupHandler struct {
	name string
}

var temp = []string{}
var temp1 = []string{}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}
func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		// fmt.Printf("%s Message topic:%q partition:%d offset:%d  value:%s\n", h.name, msg.Topic, msg.Partition, msg.Offset, string(msg.Value))

		// 先提交业务，然后再提交offset

		if msg.Topic == "diSuTestProducer-HBase" {
			// fmt.Printf("%s Message topic:%q partition:%d offset:%d  value:%s\n", h.name, msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
			temp = append(temp, string(msg.Value))
			// 由于taos时间戳不能重复插入，所以分一次插入，要是后面改为批量插入，则更改n的基数
			// 列如：1插入一条，10 插入10条，依次类推
			if len(temp) == 1 {
				tbatter.InsBatteryHistoryData(temp)
				temp = []string{}
			}
		} else if msg.Topic == "diSuTestProducer-location" {
			temp1 = append(temp1, string(msg.Value))
			// 由于taos时间戳不能重复插入，所以分一次插入，要是后面改为批量插入，则更改n的基数
			// 列如：1插入一条，10 插入10条，依次类推
			if len(temp1) == 1 {
				tbatter.InsBatteryLocationData(temp1)
				temp1 = []string{}
			}
		}
		// 手动确认消息，// 标记，sarama会自动进行提交，默认间隔1秒
		sess.MarkMessage(msg, "")
	}
	return nil
}

func handleErrors(group *sarama.ConsumerGroup, wg *sync.WaitGroup) {
	wg.Done()
	for err := range (*group).Errors() {
		fmt.Println("ERROR", err)
	}
}

func consumeHBase(group *sarama.ConsumerGroup, wg *sync.WaitGroup, name string) {
	fmt.Println(name + "start")
	wg.Done()
	ctx := context.Background()
	for {
		topics := conf.KafkaSet.KafkaTopicHBase
		handler := consumerGroupHandler{name: name}
		err := (*group).Consume(ctx, topics, handler)
		if err != nil {
			panic(err)
		}
	}
}

func consumeLocation(group *sarama.ConsumerGroup, wg *sync.WaitGroup, name string) {
	fmt.Println(name + "start")
	wg.Done()
	ctx := context.Background()
	for {
		topics := conf.KafkaSet.KafkaTopicLocation
		handler := consumerGroupHandler{name: name}
		err := (*group).Consume(ctx, topics, handler)
		if err != nil {
			panic(err)
		}
	}
}

func Consumer() {
	var wg sync.WaitGroup
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = false
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.Retry.Max = 3
	// 禁用自动提交，改为手动
	config.Consumer.Offsets.AutoCommit.Interval = time.Second * 1 // 测试1秒自动提交
	// 一般会用OffsetNewest,新老业务，新业务不需要之前的数据
	// OffsetOldest,新老业务，新业务需要之前的数据 [ 应用场景比较少],最老的  从有记录[保存周期7天，即从7天前]从0开始消费
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Version = sarama.V0_10_2_0
	client, err := sarama.NewClient(conf.KafkaSet.KafkaHost, config)
	defer client.Close()
	if err != nil {
		panic(err)
	}
	group1, err := sarama.NewConsumerGroupFromClient("HBase", client)
	if err != nil {
		panic(err)
	}
	group2, err := sarama.NewConsumerGroupFromClient("location", client)
	if err != nil {
		panic(err)
	}
	// group3, err := sarama.NewConsumerGroupFromClient("c3", client)
	// if err != nil {
	// 	panic(err)
	// }
	defer group1.Close()
	defer group2.Close()
	// defer group3.Close()
	wg.Add(2)
	go consumeHBase(&group1, &wg, "diSuTestProducer-HBase")
	go consumeLocation(&group2, &wg, "diSuTestProducer-location")
	// go consume(&group3, &wg, "c3")
	wg.Wait()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	select {
	case <-signals:
	}
}
