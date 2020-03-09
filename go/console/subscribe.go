package console

import (
	"bytes"
	"fmt"
	"github.com/go-redis/redis/v7"
	"os"
	"os/exec"
	"visible/models"
)

func Subscribe() {
	client := redis.NewClient(&redis.Options{
		Addr:               "127.0.0.1:6379",
		Password:           "",
		DB:                 0,
	})
	pubsub := client.PSubscribe("*")
	ch := pubsub.Channel()
	for msg := range ch {
		go dealMsg(msg)
	}
	defer pubsub.Close()
}

func dealMsg(msg *redis.Message) {
	topicModel := new(models.TopicModel)
	topicSubscribeList := topicModel.GetListByChannel(msg.Channel)
	if len(topicSubscribeList) != 0 {
		for _, v := range topicSubscribeList {
			go func(v models.TopicData, payload string) {
				pwd, err := os.Getwd()
				if err != nil {
					fmt.Println(err)
				}
				filename := pwd + "/../../php/public/index.php"
				_, err = os.Stat(filename)
				if err != nil {
					fmt.Println("文件不存在")
				}
				cmd := exec.Command("php", filename, v.Class, v.Function, payload)
				var stdOut, stdErr bytes.Buffer
				cmd.Stdout = &stdOut
				cmd.Stderr = &stdErr
				if err = cmd.Run(); err != nil {
					fmt.Println(err)
				}
				fmt.Println(stdOut.String(), stdErr.String())
			}(v, msg.Payload)
		}
	}
}