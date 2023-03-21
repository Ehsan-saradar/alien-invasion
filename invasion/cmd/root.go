package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"alieninvasion"

	"github.com/spf13/cobra"
)

var (
	aliensNum int
	outputMap string
)

var rootCmd = &cobra.Command{
	Use:   "invasion [path to map file]",
	Short: "Invasion is a cli app for the alien invasion challenge",
	Long: `Invasion takes the path to a map file and the number of aliens to spawn on the map.
		and runs the main algorithm of the challenge. and in the end writes the new map to a file.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Read the map file.
		mapPath := args[0]
		mapFile, err := os.ReadFile(mapPath)
		if err != nil {
			return fmt.Errorf("failed to read map file: %w", err)
		}

		// Create a new map and unmarshal the map file into it.
		var m alieninvasion.Map
		if err := m.UnmarshalText(mapFile); err != nil {
			return fmt.Errorf("failed to unmarshal map: %w", err)
		}

		// Create a new random generator with current unix time as seed.
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

		inv := alieninvasion.NewInvasion(&m, aliensNum, rnd)
		inv.Run()

		// Write the new map to a file.
		outputFile, err := os.Create(outputMap)
		if err != nil {
			return fmt.Errorf("failed to create output file: %w", err)
		}
		defer outputFile.Close()

		// Marshal the map into a byte slice.
		marshaledMap, err := inv.Map.MarshalText()
		if err != nil {
			return fmt.Errorf("failed to marshal map: %w", err)
		}

		// Write the marshaled map to the output file.
		if _, err := outputFile.Write(marshaledMap); err != nil {
			return fmt.Errorf("failed to write to output file: %w", err)
		}

		return nil
	},
}

func init() {
	rootCmd.Flags().IntVarP(&aliensNum, "aliens-num", "n", 10, "Number of aliens to spawn")
	rootCmd.Flags().StringVarP(&outputMap, "output-map", "o", "output-map.txt", "Path to the output map file")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
