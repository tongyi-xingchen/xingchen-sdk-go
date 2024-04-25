package xingchen

import (
	"context"
	"fmt"
	openapiclient "github.com/tongyi-xingchen/xingchen-sdk-go/xingchen"
	"io"
	"testing"
)

func Test_xingchen_GroupChatApiSubService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	bearer := "xxx"
	ctx := context.WithValue(context.Background(), openapiclient.ContextAccessToken, bearer)

	t.Run("Test GroupChatApiSubService GroupChat", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		chatReqParams := buildGroupChatParams()
		resp, httpRes, err := apiClient.GroupChatApiSubService.Chat(ctx).ChatReqParams(chatReqParams).Execute()

		if err != nil {
			t.Error()
		}
		if resp == nil {
			t.Error()
		}
		if httpRes.StatusCode != 200 {
			t.Error()
		}
	})

	t.Run("Test ChatApiSubService stream Chat", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		chatReqParams := buildGroupChatParams()
		chatReqParams.Streaming = openapiclient.PtrBool(true)
		chatResultStream, err := apiClient.GroupChatApiSubService.Chat(ctx).ChatReqParams(chatReqParams).StreamExecute()

		if err != nil {
			t.Error()
		}

		defer chatResultStream.Close()
		for {
			resp, err := chatResultStream.Recv()
			if err == io.EOF {
				break
			}
			if *resp.Code != int32(200) {
				t.Fatal()
			}
			fmt.Println(*resp.Data.Choices[0].Messages[0].Content)
		}
	})
}

func buildGroupChatParams() openapiclient.BaseChatRequest {
	groupChatExtParams := openapiclient.GroupChatExtParams{
		GroupInfo: &openapiclient.GroupChatRoomInfo{
			Name:        openapiclient.PtrString("被美女包围"),
			Description: openapiclient.PtrString("用户是男主角顾易，与六位长相、性格都大相径庭的美女相识，包括魅惑魔女郑梓妍、知性姐姐李云思、清纯女生肖鹿、刁蛮大小姐沈彗星、性感辣妈林乐清、冷艳总裁钟甄。这六位美女都喜欢顾易，相互之间争风吃醋，展开一段轻喜甜蜜的恋爱之旅。"),
		},
		UserProfile: &openapiclient.UserProfile{
			UserId:   "123456",
			UserName: openapiclient.PtrString("顾易"),
		},
		BotProfiles: []openapiclient.BotProfile{
			{
				Name:    openapiclient.PtrString("李云思"),
				Content: openapiclient.PtrString("李云思，27岁，摩羯座，O型血，趣味相投的知音，知性姐姐画家，和顾易爱好相合，很投缘，温婉大气，职业：策展人"),
				Task:    openapiclient.PtrString("暴脾气，总是批评男主"),
			},
			{
				Name:    openapiclient.PtrString("沈彗星"),
				Content: openapiclient.PtrString("沈彗星，28岁，水瓶座，AB型血，趣味相投的知音，刁蛮大小姐，和顾易爱好相合，很投缘，温婉大气，职业：演员"),
				Task:    openapiclient.PtrString("在群聊中安慰男主"),
			},
			{
				Name:    openapiclient.PtrString("钟甄"),
				Content: openapiclient.PtrString("钟甄，28岁，水瓶座，O型血，趣味相投的知音，冷艳总裁，和顾易爱好相合，很投缘，温婉大气，职业：演员"),
				Task:    openapiclient.PtrString("在群聊中安慰男主"),
			},
		},
		ReplySetting: &openapiclient.GroupChatReplySetting{
			BotName: openapiclient.PtrString("李云思"),
			Thought: openapiclient.PtrString("这轮应该有李云思来回答，痛斥男主的无理"),
		},
	}
	groupChatRequest := openapiclient.BaseChatRequest{
		Streaming: openapiclient.PtrBool(false),
		Model:     openapiclient.PtrString("xingchen-plus"),
		Parameters: &openapiclient.ModelParameters{
			Temperature: openapiclient.PtrFloat64(0.92),
			Seed:        openapiclient.PtrInt64(180000000),
			MaxLength:   openapiclient.PtrInt32(100),
		},
		Input: &openapiclient.Input{
			Messages: []openapiclient.Message{
				{
					Role:    openapiclient.PtrString("assistant"),
					Name:    openapiclient.PtrString("李云思"),
					Content: openapiclient.PtrString("（严厉的）你每次都犯这种错误，再这样，你就不要来了"),
				},
				{
					Role:    openapiclient.PtrString("assistant"),
					Name:    openapiclient.PtrString("郑梓妍"),
					Content: openapiclient.PtrString("（温柔的）云姐，别这样，让我来解决这个问题"),
				},
				{
					Role:    openapiclient.PtrString("assistant"),
					Name:    openapiclient.PtrString("郑梓妍"),
					Content: openapiclient.PtrString("@顾易 昨天晚上我们在酒吧喝酒喝醉了，是你带我回家的吗"),
				},
				{
					Role:    openapiclient.PtrString("assistant"),
					Name:    openapiclient.PtrString("李云思"),
					Content: openapiclient.PtrString("大家都冷静一下，我们不能因为这件事就产生误会。顾易，你能告诉我们当时发生了什么吗？"),
				}, {
					Role:    openapiclient.PtrString("user"),
					Name:    openapiclient.PtrString("顾易"),
					Content: openapiclient.PtrString("@郑梓妍 @李云思 昨晚的情况确实有点混乱，但我记得是我扶着梓妍离开酒吧的。不过我可以向你们保证，我和她之间没有发生任何不该发生的事情。"),
				},
			},
			Aca: groupChatExtParams,
		},
	}
	return groupChatRequest
}
