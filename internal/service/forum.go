package service

import (
	"myChat/internal/domain/model"
	"myChat/internal/domain/repository"
	"myChat/pkg/apperrors"
	"myChat/pkg/utils"
	"time"
)

type ForumService struct {
	tRepo repository.ThreadRepository
	pRepo repository.PostRepository
	uRepo repository.UserRepository
}

func NewForumService(t repository.ThreadRepository, p repository.PostRepository, u repository.UserRepository) *ForumService {
	return &ForumService{
		tRepo: t,
		pRepo: p,
		uRepo: u,
	}
}

type ThreadList []struct {
	Topic      string
	UserName   string
	CreatedAt  string
	NumReplies int
	Uuid       string
}

func (fs *ForumService) ReadThreadList() (ThreadList, error) {

	// create template data
	// send to index page template.

	var data ThreadList
	threads, err := fs.tRepo.FindAll()
	if err != nil {
		err = apperrors.ReadThreadFailed.Wrap(err, "failed to read all threads")
		return data, err
	}

	for _, thread := range threads {
		// Get user who wrote this thread.
		user, err := fs.uRepo.FindById(thread.UserId)
		if err != nil {
			err = apperrors.NoUserFound.Wrap(err, "failed to found user by id")
			return data, err
		}

		// Get post number that is replied on this thread.
		replies, err := fs.tRepo.CountPostNum(thread.Id)
		if err != nil {
			err = apperrors.CountRepliesFailed.Wrap(err, "fail to count replies on thread")
			return data, err
		}

		// Generate data format for template
		data = append(data, struct {
			Topic      string
			UserName   string
			CreatedAt  string
			NumReplies int
			Uuid       string
		}{
			Topic:      thread.Topic,
			UserName:   user.Name,
			CreatedAt:  thread.CreatedAtStr(),
			NumReplies: replies,
			Uuid:       thread.Uuid,
		})
	}

	return data, nil
}

func (fs *ForumService) CreateThread(userId int, topic string) error {
	user, err := fs.uRepo.FindById(userId)
	if err != nil {
		err = apperrors.NoUserFound.Wrap(err, "failed to found user by id")
		return err
	}

	thread := model.Thread{
		Uuid:      utils.CreateUUID(),
		Topic:     topic,
		UserId:    user.Id,
		CreatedAt: time.Now(),
	}
	if err := fs.tRepo.Save(thread); err != nil {
		err = apperrors.CreateThreadFailed.Wrap(err, "failed to save thread data")
		return err
	}
	return nil
}

func (fs *ForumService) CreatePost(userId int, body string, threadUuid string) error {
	user, err := fs.uRepo.FindById(userId)
	if err != nil {
		err = apperrors.NoUserFound.Wrap(err, "failed to found user by id")
		return err
	}

	thread, err := fs.tRepo.FindByUuid(threadUuid)
	if err != nil {
		err = apperrors.ReadThreadFailed.Wrap(err, "failed to read thread by uuid")
		return err
	}

	post := model.Post{
		Uuid:      utils.CreateUUID(),
		Body:      body,
		UserId:    user.Id,
		ThreadId:  thread.Id,
		CreatedAt: time.Now(),
	}
	if err := fs.pRepo.Save(post); err != nil {
		apperrors.CreatePostFailed.Wrap(err, "failed to save post data")
		return err
	}
	return nil
}

type ThreadDetail struct {
	Topic     string
	UserName  string
	CreatedAt string
	Uuid      string
	Posts     []struct {
		Body      string
		UserName  string
		CreatedAt string
	}
}

func (fs *ForumService) ReadThreadDetail(uuid string) (ThreadDetail, error) {
	// Create template data.
	var data ThreadDetail

	thread, err := fs.tRepo.FindByUuid(uuid)
	if err != nil {
		err = apperrors.ReadThreadFailed.Wrap(err, "failed to read thread by uuid")
		return data, err
	}

	user, err := fs.uRepo.FindById(thread.UserId)
	if err != nil {
		err = apperrors.NoUserFound.Wrap(err, "failed to get user by id")
		return data, err
	}

	posts, err := fs.pRepo.FindByThreadId(thread.Id)
	if err != nil {
		err = apperrors.ReadPostFailed.Wrap(err, "failed to read post by thread_id")
		return data, err
	}
	data.Posts = make([]struct {
		Body      string
		UserName  string
		CreatedAt string
	}, len(posts))

	data.Topic = thread.Topic
	data.UserName = user.Name
	data.CreatedAt = thread.CreatedAtStr()
	data.Uuid = thread.Uuid

	// Insert post data into thread data.
	for i, post := range posts {
		user, err := fs.uRepo.FindById(post.UserId)
		if err != nil {
			err = apperrors.NoUserFound.Wrap(err, "failed to get user by post")
			return data, err
		}

		data.Posts[i].Body = post.Body
		data.Posts[i].UserName = user.Name
		data.Posts[i].CreatedAt = post.CreatedAtStr()
	}

	return data, nil
}
