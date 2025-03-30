package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"Sranjan0208/go-todo/board"
	"Sranjan0208/go-todo/cli"
	"Sranjan0208/go-todo/utils"
)

var boardInstance = &board.Board{Tasks: utils.LoadTasks()}

func main() {
	var rootCmd = &cobra.Command{Use: "kanban"}

	// Pass boardInstance to CLI commands
	cli.SetBoardInstance(boardInstance)

	// Register CLI commands
	rootCmd.AddCommand(cli.AddCmd)
	rootCmd.AddCommand(cli.ListCmd)
	rootCmd.AddCommand(cli.UpdateCmd)
	rootCmd.AddCommand(cli.DeleteCmd)
	rootCmd.AddCommand(cli.MoveCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
