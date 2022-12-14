package iox

import "os"

func WriteAllText(file, content string) error {
	fo, err := os.Create(file)
	if err != nil {
		return err
	}
	defer fo.Close()
	_, err = fo.WriteString(content)
	return err
}

func WriteAllBytes(file string, content []byte) error {
	return WriteAllText(file, string(content))
}

func WriteAllLines(file string, content []string) error {
	fo, err := os.Create(file)
	if err != nil {
		return err
	}
	defer fo.Close()
	for _, line := range content {
		_, err = fo.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
