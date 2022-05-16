package main

import (
	"bufio"
	"bytes"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
)

func deleteRecordsCmd() *cobra.Command {
	var (
		filePath string
	)
	cmd := &cobra.Command{
		Use:   `delete-records`,
		Short: `позволяет массово удалять записи с числовыми айдишниками, хранимыми в файле построчно`,
		Run: func(cmd *cobra.Command, args []string) {
			doDeleteRecordsByIds(zoneName, filePath)
		},
	}
	cmd.PersistentFlags().StringVar(&filePath, `file-path`, ``, `файл со списком айдишников`)
	return cmd
}

func deleteRecord() *cobra.Command {
	var (
		id int
	)
	cmd := &cobra.Command{
		Use:   `delete-record`,
		Short: `удалять одну запись`,
		Run: func(cmd *cobra.Command, args []string) {
			if _, err := apiClient.DeleteRecord(zoneName, id); err != nil {
				logrus.Fatalln(err)
			} else {
				logrus.Infof(`id %d deleted`, id)
			}
		},
	}
	cmd.PersistentFlags().IntVar(&id, `id`, 0, `айдишник записи для удаления`)
	return cmd
}

func doDeleteRecordsByIds(zoneName string, filePath string) {

	data, err := os.ReadFile(filePath)
	if err != nil {
		logrus.Fatalln(`file error: %s`, err)
	}
	scanner := bufio.NewScanner(bytes.NewBuffer(data))
	var ids []int64
	for scanner.Scan() {
		sid := strings.Split(scanner.Text(), " ")[0]
		v, err := strconv.ParseInt(sid, 10, 64)
		if err != nil {
			logrus.Fatalln(`value %s error: %s`, sid, err.Error())
		}
		ids = append(ids, v)
	}

	for _, id := range ids {
		if _, err := apiClient.DeleteRecord(zoneName, int(id)); err != nil {
			logrus.Errorln(id, err)
		} else {
			logrus.Infof(`id %d deleted`, id)
		}
	}
}
