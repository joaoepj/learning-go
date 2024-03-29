// learning-go is a repository showing how to apply development
// best practices in your code while learning a programming language
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joaoepj/learning-go/lg_algorithms"
	"github.com/joaoepj/learning-go/lg_misc"
	"github.com/joaoepj/learning-go/lg_problems"
	"github.com/urfave/cli/v2"
)

var (
	LearningGoGitHash   = "None"
	LearningGoGitBranch = "None"
)

func main() {

	log.SetFlags(log.Lshortfile)

	app := &cli.App{
		Name:  "learning-go",
		Usage: "a repository to learn the go programming language.",
		UsageText: `learning-go command [command options]
learning-go [global options]`,
		HideVersion: false,
		Version:     LearningGoGitBranch + "-" + LearningGoGitHash,
		Flags:       []cli.Flag{},
		Action: func(c *cli.Context) error {
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "algorithms",
				Aliases: []string{"a"},
				Usage:   "The Algorithm Design Manual",
				Subcommands: []*cli.Command{
					lg_algorithms.SCFibonacci,
				},
			},
			{
				Name:    "problems",
				Aliases: []string{"p"},
				Usage:   "Some HackerRank problem solutions implemented in golang.",
				Subcommands: []*cli.Command{
					{
						Name:    "arraymanipulation",
						Aliases: []string{"am"},
						Usage:   "Array Manipulation Problem",
						Action: func(c *cli.Context) error {
							for _, arr := range lg_problems.ArrayManipulation_arr {
								lg_problems.ArrayManipulation(10, arr)
							}
							return nil
						},
					},

					{
						Name:    "bigsum",
						Aliases: []string{"bs"},
						Usage:   "A Very Big Sum Problem",
						Action: func(c *cli.Context) error {
							lg_problems.AVeryBigSum([]int64{1.0, 2.0, 3.0})
							return nil
						},
					},
					{
						Name:    "catmouse",
						Aliases: []string{"cm"},
						Usage:   "Cats and a Mouse Problem",
						Action: func(c *cli.Context) error {
							lg_problems.CatAndMouse(5, 9, 13)
							return nil
						},
					},
					{
						Name:    "checkmagazine",
						Aliases: []string{"cm2"},
						Usage:   "Check Magazine Problem",
						Action: func(c *cli.Context) error {
							lg_problems.CheckMagazine(lg_problems.Magazine, lg_problems.Note)
							return nil
						},
					},

					{
						Name:    "climbinglearderboard",
						Aliases: []string{"cl"},
						Usage:   "Climbing the Leaderboard Problem",
						Action: func(c *cli.Context) error {
							lg_problems.ClimbingLeaderboard2([]int32{100, 100, 50, 40, 40, 20, 10}, []int32{5, 25, 50, 120})
							return nil
						},
					},
					{
						Name:    "comparethetriplets",
						Aliases: []string{"ctt"},
						Usage:   "Compare The Triplets Problem",
						Action: func(c *cli.Context) error {
							t1 := []int32{1, 2, 3}
							t2 := []int32{1, 2, 3}
							lg_problems.CompareTriplets(t1, t2)
							return nil
						},
					},
					{
						Name:    "designerpdfviewer",
						Aliases: []string{"dpv"},
						Usage:   "Designer Pdf Viewer Problem",
						Action: func(c *cli.Context) error {

							lg_problems.DesignerPdfViewer(lg_problems.Alphabet_height, "reality")

							return nil
						},
					},
					{
						Name:    "formingmagicsquare",
						Aliases: []string{"fms"},
						Usage:   "Forming A Magic Square Problem",
						Action: func(c *cli.Context) error {
							lg_problems.FormingMagicSquare(lg_problems.TheSquare)
							return nil
						},
					},
					{
						Name:    "hourglasssum",
						Aliases: []string{"hgs"},
						Usage:   "Hourglass Sum Problem",
						Action: func(c *cli.Context) error {
							lg_problems.HourglassSum(lg_problems.HourglassSum_arr)
							return nil
						},
					},
					{
						Name:    "hurdlerace",
						Aliases: []string{"hr"},
						Usage:   "Hurdle Race Problem",
						Action: func(c *cli.Context) error {
							lg_problems.Staircase(5)
							return nil
						},
					},
					{
						Name:    "letfrotation",
						Aliases: []string{"lr"},
						Usage:   "Left Rotation Problem",
						Action: func(c *cli.Context) error {
							lg_problems.RotLeft([]int32{1, 2, 3, 4, 5}, 3)
							return nil
						},
					},
					{
						Name:    "migratorybirds",
						Aliases: []string{"mb"},
						Usage:   "Migratory Birds Problem",
						Action: func(c *cli.Context) error {
							lg_problems.MigratoryBirds(lg_problems.MigratoryBirds_arr)
							return nil
						},
					},
					{
						Name:    "minimumbribes",
						Aliases: []string{"mb2"},
						Usage:   "Minimun Bribes Problem",
						Action: func(c *cli.Context) error {
							for _, arr := range lg_problems.MinimumBribes_arr {
								lg_problems.MinimumBribes(arr)
							}
							return nil
						},
					},
					{
						Name:    "minimumswaps2",
						Aliases: []string{"ms2"},
						Usage:   "Minimun Swaps 2 Problem",
						Action: func(c *cli.Context) error {
							for _, arr := range lg_problems.MinimumSwaps2_arr {
								lg_problems.MinimumSwaps2(arr)
							}
							return nil
						},
					},

					{
						Name:    "minmaxsum",
						Aliases: []string{"mms"},
						Usage:   "Min-Max Sum Problem",
						Action: func(c *cli.Context) error {
							arr := []int32{2, 8, 5, 9, 2, 5, 7, 1}
							lg_problems.MinMaxSum(arr)
							return nil
						},
					},
					{
						Name:    "missinginteger",
						Aliases: []string{"mi"},
						Usage:   "Missing Integer Problem",
						Action: func(c *cli.Context) error {
							lg_problems.TestMissingInteger()
							return nil
						},
					},
					{
						Name:    "staircase",
						Aliases: []string{"s"},
						Usage:   "Staircase Problem",
						Action: func(c *cli.Context) error {
							lg_problems.Staircase(5)
							return nil
						},
					},
					{
						Name:    "twostrings",
						Aliases: []string{"ts"},
						Usage:   "Two String Problem",
						Action: func(c *cli.Context) error {
							lg_problems.TwoStrings(lg_problems.TSMajor, lg_problems.TSMinor)
							return nil
						},
					},
					lg_problems.SCCountSwaps,
					lg_problems.SCTwoSum,
					lg_problems.SCPickingNumbers,
					lg_problems.SCAngryProfessor,
					lg_problems.SCBeautifulDays,
					lg_problems.SCUtopianTree,
					lg_problems.SCViralAdvertisement,
				}, // Closing problems Subcommands
			},

			{
				Name:    "misc",
				Aliases: []string{"m"},
				Usage:   "Miscelaneous code",
				Subcommands: []*cli.Command{
					{
						Name:    "asn1",
						Aliases: []string{"as"},
						Usage:   "Test encoding/decoding time with ASN.1 library",
						Action: func(c *cli.Context) error {
							lg_misc.Asn1()
							return nil
						},
					},
					{
						Name:    "mediansortedarrays",
						Aliases: []string{"msa"},
						Usage:   "Find Median Sorted Arrays",
						Action: func(c *cli.Context) error {
							x := []int{1, 2}
							y := []int{3, 4}
							fmt.Println("v1: ", lg_misc.FindMedianSortedArraysV1(x, y))
							fmt.Println("v2: ", lg_misc.FindMedianSortedArraysV2(x, y))
							fmt.Println("v3: ", lg_misc.FindMedianSortedArraysV3(x, y))
							return nil
						},
					},
					{
						Name:    "snmpget",
						Aliases: []string{"sg"},
						Usage:   "snmpget",
						Action: func(c *cli.Context) error {
							lg_misc.SnmpGet()
							return nil
						},
					},
					{
						Name:    "snmpwalk",
						Aliases: []string{"sw"},
						Usage:   "snmpwalt",
						Action: func(c *cli.Context) error {
							lg_misc.SnmpWalk()
							return nil
						},
					},
					{
						Name:    "testfdlimit",
						Aliases: []string{"tfl"},
						Usage:   "Testing the Limit for File Descriptors",
						Action: func(c *cli.Context) error {
							lg_misc.TestFdLimit()
							return nil
						},
					},

					lg_misc.SCRandomString,
					lg_misc.SCRandomIntegers,
					lg_misc.SCRecursive,
					lg_misc.SCSorting,
					lg_misc.SCRemoveVowels,
				}, // Closing Misc Subcommands
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
