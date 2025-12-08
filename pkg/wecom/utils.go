package wecom

import "strings"

func dedupStrings(items []string) []string {
	seen := map[string]struct{}{}
	var out []string
	for _, item := range items {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		if _, ok := seen[item]; ok {
			continue
		}
		seen[item] = struct{}{}
		out = append(out, item)
	}
	return out
}

// buildUserTokens 生成 <@userid> 便于 Markdown 可视化提醒
func buildUserTokens(userIDs []string) []string {
	userIDs = dedupStrings(userIDs)
	if len(userIDs) == 0 {
		return nil
	}
	var tokens []string
	for _, id := range userIDs {
		if strings.HasPrefix(id, "<@") {
			tokens = append(tokens, id)
			continue
		}
		tokens = append(tokens, "<@"+id+">")
	}
	return tokens
}
