package ini

import "strings"

type IniConfig map[string]map[string]string

func ParseLines(s []string) IniConfig {
	section := ""
	config := make(IniConfig)
	config[""] = make(map[string]string)
	for _, line := range s {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if line[0] == '[' {
			section = line[1 : len(line)-1]
			config[section] = make(map[string]string)
		} else {
			kv := strings.SplitN(line, "=", 2)
			if len(kv) != 2 {
				continue
			}
			key := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			config[section][key] = value
		}
	}
	return config
}

func ParseString(s string) IniConfig {
	return ParseLines(strings.Split(s, "\n"))
}

func (c IniConfig) String() string {
	var s strings.Builder
	dft, ok := c[""]
	if ok {
		for k, v := range dft {
			s.WriteString(k)
			s.WriteString("=")
			s.WriteString(v)
			s.WriteString("\n")
		}
	}

	for section, kv := range c {
		if section == "" {
			continue
		}
		s.WriteString("[")
		s.WriteString(section)
		s.WriteString("]\n")
		for key, value := range kv {
			s.WriteString(key)
			s.WriteString("=")
			s.WriteString(value)
			s.WriteString("\n")
		}
	}
	return s.String()
}
