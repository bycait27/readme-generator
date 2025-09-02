package prompts

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

// common utility functions used across all prompts

// promptYesNo asks a yes/no questions
func promptYesNo(label string) (bool, error) {
	prompt := promptui.Select{
		Label: label,
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		return false, err
	}
	return result == "Yes", nil
}

// promptRequiredText prompts for required text with validation
func promptRequiredText(label string, minLength, maxLength int) (string, error) {
	prompt := promptui.Prompt{
		Label: label,
		Validate: func(input string) error {
			trimmed := strings.TrimSpace(input)
			if len(trimmed) < minLength {
				return fmt.Errorf("input must be at least %d characters", minLength)
			}
			if len(input) > maxLength {
				return fmt.Errorf("input must be at most %d characters", maxLength)
			}
			return nil
		},
	}
	return prompt.Run()
}

// promptOptionalText prompts for optional text
func promptOptionalText(label string, maxLength int) (string, error) {
	wantText, err := promptYesNo(fmt.Sprintf("Do you want to add %s?", strings.ToLower(label)))
	if err != nil {
		return "", err
	}

	if !wantText {
		return "", nil
	}

	return promptRequiredText(label, 1, maxLength)
}

// promptRequiredInt prompts for required int with validation
func promptRequiredInt(label string, minValue, maxValue int) (int, error) {
	prompt := promptui.Prompt{
		Label: label,
		Validate: func(input string) error {
			trimmed := strings.TrimSpace(input)
			if trimmed == "" {
				return fmt.Errorf("input is required")
			}
			val, err := strconv.Atoi(trimmed)
			if err != nil {
				return fmt.Errorf("please enter a valid integer")
			}
			if val < minValue {
				return fmt.Errorf("input must be at least %d", minValue)
			}
			if val > maxValue {
				return fmt.Errorf("input must be at most %d", maxValue)
			}
			return nil
		},
	}
	result, err := prompt.Run()
	if err != nil {
		return 0, err
	}
	val, err := strconv.Atoi(strings.TrimSpace(result))
	if err != nil {
		return 0, fmt.Errorf("failed to parse integer: %w", err)
	}
	return val, nil
}

// promptOptionalInt prompts for optional int
func promptOptionalInt(label string, minValue, maxValue int) (int, error) {
	wantInt, err := promptYesNo(fmt.Sprintf("Do you want to add %s?", strings.ToLower(label)))
	if err != nil {
		return 0, err
	}

	if !wantInt {
		return 0, nil
	}

	return promptRequiredInt(label, minValue, maxValue)
}

// promptOptionalStringPointer prompts for optional text that returns *string
func promptOptionalStringPointer(label string, maxLength int) (*string, error) {
	wantText, err := promptYesNo(fmt.Sprintf("Do you want to add %s?", strings.ToLower(label)))
	if err != nil {
		return nil, nil
	}

	if !wantText {
		return nil, nil
	}

	text, err := promptRequiredText(label, 1, maxLength)
	if err != nil {
		return nil, err
	}

	return &text, nil
}

// promptStringList collects a list of strings
func promptStringList(itemName string, required bool) ([]string, error) {
	wantItems, err := promptYesNo(fmt.Sprintf("Do you want to add %s?", itemName))
	if err != nil {
		return nil, err
	}

	if !wantItems {
		if required {
			return nil, fmt.Errorf("%s are required", itemName)
		}
		return nil, nil
	}

	var items []string

	for {
		prompt := promptui.Prompt{
			Label: fmt.Sprintf("Enter %s (or press Enter to finish)", itemName),
			Validate: func(input string) error {
				trimmed := strings.TrimSpace(input)
				if trimmed == "" && len(items) == 0 && required {
					return fmt.Errorf("at least one %s is required", strings.TrimSuffix(itemName, "s"))
				}
				return nil
			},
		}

		item, err := prompt.Run()
		if err != nil {
			return nil, err
		}

		item = strings.TrimSpace(item)
		if item == "" {
			break
		}

		items = append(items, item)
	}

	return items, nil
}

// promptFromOptions prompts user to select from predefined options
func promptFromOptions(label string, options []string) (string, error) {
	prompt := promptui.Select{
		Label: label,
		Items: options,
	}
	_, result, err := prompt.Run()
	return result, err
}

// promptRequiredBoolean prompts the user to provide a boolean response
func promptRequiredBoolean(label string) (bool, error) {
	return promptYesNo(label)
}

// promptURL prompts for a URL with validation
func promptURL(label string, required bool) (string, error) {
	prompt := promptui.Prompt{
		Label: label,
		Validate: func(input string) error {
			trimmed := strings.TrimSpace(input)
			if !required && trimmed == "" {
				return nil
			}
			// URL validation
			if trimmed != "" && !strings.HasPrefix(trimmed, "http://") && !strings.HasPrefix(trimmed, "https://") {
				return errors.New("URL must start with http:// or https://")
			}
			return nil
		},
	}
	return prompt.Run()
}

// promptEmail prompts for email with basic validation
func promptEmail(label string) (string, error) {
	prompt := promptui.Prompt{
		Label: label,
		Validate: func(input string) error {
			trimmed := strings.TrimSpace(input)
			if trimmed == "" {
				return errors.New("email is required")
			}
			if !strings.Contains(trimmed, "@") {
				return errors.New("please enter a valid email address")
			}
			return nil
		},
	}
	return prompt.Run()
}
