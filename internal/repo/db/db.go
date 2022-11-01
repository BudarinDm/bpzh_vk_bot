package db

import (
	"bpzh_vk_bot/internal/config"
	"cloud.google.com/go/firestore"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/api/option"
)

type Repo struct {
	DB *pgxpool.Pool
	FS *firestore.Client
}

func NewRepo(FS *firestore.Client) *Repo {
	return &Repo{
		FS: FS,
	}
}

func CreateFSConnections(cfg *config.DBConfig) (*firestore.Client, error) {
	//err := ioutil.WriteFile("fs.json", []byte(cfg.FSConf), 0777)
	//if err != nil {
	//	fmt.Println(err)
	//}

	options := option.WithCredentialsFile("./bot_fs.json")
	client, err := firestore.NewClient(context.Background(), "bpzh-info", options)
	if err != nil {
		return nil, err
	}
	return client, nil
}
