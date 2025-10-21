package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "No DNS .env"
	}

	db, err := openDB(dsn)
	if err != nil {
		log.Fatalf("openDB error: %v", err)
	}
	defer db.Close()

	repo := NewRepo(db)

	// 1) Вставим пару задач
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	titles := []string{"Сделать ПЗ №5", "Купить кофе", "Проверить отчёты"}
	for _, title := range titles {
		id, err := repo.CreateTask(ctx, title)
		if err != nil {
			log.Fatalf("CreateTask error: %v", err)
		}
		log.Printf("Inserted task id=%d (%s)", id, title)
	}

	ctxMany, cancelMany := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelMany()

	bulkTitles := []string{"Выаолнение задания 1", "Выаолнение задания 2", "Выаолнение задания 3"}
	err = repo.CreateMany(ctxMany, bulkTitles)
	if err != nil {
		log.Fatalf("CreateMany error: %v", err)
	}
	log.Println("Массовая вставка завершена")

	// 2) Прочитаем список задач
	ctxList, cancelList := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelList()

	tasks, err := repo.ListTasks(ctxList)
	if err != nil {
		log.Fatalf("ListTasks error: %v", err)
	}

	// 3) Напечатаем
	fmt.Println("=== Tasks ===")
	for _, t := range tasks {
		fmt.Printf("#%d | %-24s | done=%-5v | %s\n",
			t.ID, t.Title, t.Done, t.CreatedAt.Format(time.RFC3339))
	}

	ctxDone, cancelDone := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelDone()

	notDoneTasks, err := repo.ListDone(ctxDone, false)
	if err != nil {
		log.Fatalf("ListDone error: %v", err)
	}

	fmt.Println("=== Not Done Tasks ===")
	for _, t := range notDoneTasks {
		fmt.Printf("#%d | %-24s | done=%-5v | %s\n",
			t.ID, t.Title, t.Done, t.CreatedAt.Format(time.RFC3339))
	}

	ctxFind, cancelFind := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFind()

	task, err := repo.FindByID(ctxFind, 1)
	if err != nil {
		log.Fatalf("FindByID error: %v", err)
	}

	fmt.Println("=== Task with ID=1 ===")
	fmt.Printf("#%d | %-24s | done=%-5v | %s\n",
		task.ID, task.Title, task.Done, task.CreatedAt.Format(time.RFC3339))
}
