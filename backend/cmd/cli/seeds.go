package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"math/rand"
	"strings"
)
var (
	adjectives = []string{
		"premium", "innovative", "durable", "lightweight", "ergonomic", "versatile", 
		"eco-friendly", "high-performance", "sleek", "robust", "compact", "stylish",
		"advanced", "reliable", "efficient", "modern", "professional", "cutting-edge",
		"superior", "enhanced", "smart", "wireless", "waterproof", "portable",
	}
	
	materials = []string{
		"aluminum", "carbon fiber", "stainless steel", "titanium", "polymer", 
		"ceramic", "bamboo", "recycled plastic", "tempered glass", "silicone",
		"leather", "fabric", "wood", "composite", "metal alloy",
	}
	
	features = []string{
		"with built-in LED indicators", "featuring wireless connectivity", 
		"equipped with advanced sensors", "including premium accessories",
		"with touchscreen interface", "featuring voice control", 
		"including smartphone integration", "with automatic adjustment",
		"featuring rapid charging", "with cloud synchronization",
		"including GPS tracking", "with anti-theft protection",
		"featuring noise cancellation", "with gesture recognition",
	}
	
	benefits = []string{
		"Perfect for professionals and enthusiasts alike",
		"Designed to enhance your daily productivity",
		"Built to withstand demanding conditions",
		"Engineered for maximum comfort and efficiency",
		"Crafted with attention to every detail",
		"Optimized for both indoor and outdoor use",
		"Suitable for beginners and experts",
		"Guaranteed to exceed your expectations",
		"Backed by industry-leading warranty",
		"Trusted by thousands of satisfied customers",
	}
	
	categories = []string{
		"Electronics", "Home & Garden", "Sports & Outdoors", "Automotive", 
		"Health & Wellness", "Office Supplies", "Kitchen & Dining", "Tools & Hardware",
		"Fashion & Accessories", "Technology", "Fitness Equipment", "Travel Gear",
	}
)

func generateRandomDescription(productNum int) string {
	var desc strings.Builder
	
	// Category and basic description
	category := categories[rand.Intn(len(categories))]
	adj1 := adjectives[rand.Intn(len(adjectives))]
	adj2 := adjectives[rand.Intn(len(adjectives))]
	material := materials[rand.Intn(len(materials))]
	
	desc.WriteString(fmt.Sprintf("This %s %s product represents the perfect blend of %s design and %s construction. ", 
		adj1, category, adj2, material))
	
	// Features
	numFeatures := rand.Intn(3) + 2 // 2-4 features
	selectedFeatures := make([]string, 0, numFeatures)
	usedFeatures := make(map[int]bool)
	
	for len(selectedFeatures) < numFeatures {
		idx := rand.Intn(len(features))
		if !usedFeatures[idx] {
			selectedFeatures = append(selectedFeatures, features[idx])
			usedFeatures[idx] = true
		}
	}
	
	desc.WriteString("Our engineering team has incorporated cutting-edge technology ")
	for i, feature := range selectedFeatures {
		if i == 0 {
			desc.WriteString(feature)
		} else if i == len(selectedFeatures)-1 {
			desc.WriteString(", and " + feature)
		} else {
			desc.WriteString(", " + feature)
		}
	}
	desc.WriteString(". ")
	
	// Benefits and specifications
	benefit := benefits[rand.Intn(len(benefits))]
	desc.WriteString(benefit + " ")
	
	// Technical specifications
	weight := rand.Float32()*5 + 0.5 // 0.5-5.5 lbs
	dimensions := fmt.Sprintf("%.1fx%.1fx%.1f", 
		rand.Float32()*20+5,  // 5-25 inches
		rand.Float32()*15+3,  // 3-18 inches  
		rand.Float32()*10+1)  // 1-11 inches
	
	desc.WriteString(fmt.Sprintf("With precise dimensions of %s inches and weighing only %.1f pounds, ", 
		dimensions, weight))
	
	// Final selling points
	warranty := []string{"2-year", "3-year", "5-year", "lifetime"}
	selectedWarranty := warranty[rand.Intn(len(warranty))]
	
	desc.WriteString(fmt.Sprintf("this product offers exceptional value and reliability. Comes with our comprehensive %s warranty and 30-day money-back guarantee. ", selectedWarranty))
	
	// Model number and compliance
	modelNum := fmt.Sprintf("MD-%d-%s", productNum, strings.ToUpper(string(rune('A'+rand.Intn(26)))))
	desc.WriteString(fmt.Sprintf("Model: %s. Meets all industry standards and certifications.", modelNum))
	
	return desc.String()
}


func newSeedsCmd(pc productCatalog) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "seeds",
		Short: "populate the database with initial data",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			for i := 0; i < 1000; i++ {
				pID := fmt.Sprintf("product-%d", i)
				name := fmt.Sprintf("Product %d", i)
				desc := generateRandomDescription(i + 1)

				if err := pc.Add(ctx, pID, name, desc, 100, "USD"); err != nil {
					return err
				}

			}

			return nil
		},
	}

	return cmd
}
