package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"strings"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

// addLabel adds a text label to the image at the specified coordinates with the given color
func addLabel(img *image.RGBA, x, y int, label string, col color.Color) {
	point := fixed.Point26_6{fixed.I(x), fixed.I(y)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}

// fillBackground fills the entire image with the specified background color
func fillBackground(img *image.RGBA, col color.Color) {
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			img.Set(x, y, col)
		}
	}
}

// colorToString converts a color.Color to its string representation in RGBA format
func colorToString(col color.Color) string {
	r, g, b, a := col.RGBA()
	return fmt.Sprintf("rgba(%d, %d, %d, %d)", r>>8, g>>8, b>>8, a>>8)
}

// colorToName maps a color.Color to its common name using a predefined color map
func colorToName(col color.Color) string {
	colorMap := map[string]string{
		"rgba(255, 0, 0, 255)":     "Red",
		"rgba(0, 255, 0, 255)":     "Green",
		"rgba(0, 0, 255, 255)":     "Blue",
		"rgba(0, 0, 0, 255)":       "Black",
		"rgba(255, 255, 255, 255)": "White",
		"rgba(255, 255, 0, 255)":   "Yellow",
		"rgba(255, 0, 255, 255)":   "Magenta",
		"rgba(0, 255, 255, 255)":   "Cyan",
		"rgba(128, 128, 128, 255)": "Gray",
	}
	return colorMap[colorToString(col)]
}

func main() {
	// Create a new RGBA image with specified dimensions
	img := image.NewRGBA(image.Rect(0, 0, 300, 100))

	// Set the background color to light gray
	fillBackground(img, color.RGBA{240, 240, 240, 255})

	// Define unique colors for each label
	col := []color.RGBA{
		{255, 0, 0, 255},     // Red
		{0, 255, 0, 255},     // Green
		{0, 0, 255, 255},     // Blue
		{0, 0, 0, 255},       // Black
		{255, 255, 255, 255}, // White
		{255, 255, 0, 255},   // Yellow
		{255, 0, 255, 255},   // Magenta
		{0, 255, 255, 255},   // Cyan
		{128, 128, 128, 255}, // Gray
	}

	// Define the labels with their positions, colors, and text
	labels := []struct {
		x, y int
		col  color.Color
		text string
	}{
		{20, 30, col[0], GenText(string(make([]byte, 1)))},
		{120, 30, col[1], GenText(string(make([]byte, 2)))},
		{220, 30, col[2], GenText(string(make([]byte, 3)))},
		{20, 60, col[3], GenText(string(make([]byte, 4)))},
		{120, 60, col[4], GenText(string(make([]byte, 5)))},
		{220, 60, col[5], GenText(string(make([]byte, 6)))},
		{20, 90, col[6], GenText(string(make([]byte, 7)))},
		{120, 90, col[7], GenText(string(make([]byte, 8)))},
		{220, 90, col[8], GenText(string(make([]byte, 9)))},
	}

	// Add each label to the image
	for _, label := range labels {
		addLabel(img, label.x, label.y, label.text, label.col)
	}

	// Create a new file to save the image
	f, err := os.Create("laurie.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Encode the image to PNG format and save it to the file
	if err := png.Encode(f, img); err != nil {
		panic(err)
	}

	// Print the generated labels for reference
	fmt.Println("Generated labels:")
	for _, label := range labels {
		fmt.Printf("Text: %s, Color: %s\n", label.text, colorToName(label.col))
	}

	// Choose a random label color
	rand.Seed(time.Now().UnixNano())
	randomLabel := labels[rand.Intn(len(labels)-1)+1] // Exclude the question label
	randomColor := colorToName(randomLabel.col)
	fmt.Printf("Please choose the label with the color: %s\n", randomColor)

	// Ask user to choose a label based on its color
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter the text of the label with the specified color: ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput) // Remove newline and any surrounding whitespace

		if userInput == randomLabel.text {
			fmt.Println("Correct! You chose the right label.")
			break
		} else {
			fmt.Println("Incorrect. Please try again.")
		}
	}
}

// GenText generates a random text string of length 8 using the provided salt
func GenText(salt string) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890@$&"
	seed := int64(time.Now().UnixNano()) + int64(len(salt)) // Convert to int64
	r := rand.New(rand.NewSource(seed))
	text := ""
	for i := 0; i < 8; i++ {
		text += string(chars[r.Intn(len(chars))])
	}
	return text
}

// resources: stackoverflow.com, pkg.go.dev, github.com, openai.com
// Stackoverflow: GenText, font, addLabel, fillBackground. example: https://stackoverflow.com/questions/35781197/generating-a-random-fixed-length-byte-array-in-go
// Golang: image, color, math
// Github: GenText
// OpenAI: col, colorToString, colorToName, userInput, Comments - I should've saved the prompts
