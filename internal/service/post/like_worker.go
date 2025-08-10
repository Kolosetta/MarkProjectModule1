package post

import (
	"errors"
	"log"
)

type LikeTask struct {
	PostId int64
	UserId int64
}

var likeChan = make(chan LikeTask, 10)

// Воркер крутится в фоне, и горутина заблокирована, пока канал likeChan пуст
func StartLikeWorker(s *Service) {
	go func() {
		for task := range likeChan {
			err := s.LikePost(task.UserId, task.PostId)
			if err != nil {
				log.Println("Ошибка при добавлении лайка:", err)
			}
		}
	}()
}

func AddLikeAsync(postId, userId int64) error {
	likeChan <- LikeTask{postId, userId}

	select {
	case likeChan <- LikeTask{postId, userId}:
		return nil
	default:
		//Очередь переполнена
		return errors.New("queue is full")
	}
}
