package components

import (
	"fmt"
	"log"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/JohannesKaufmann/html-to-markdown/plugin"
	"github.com/charmbracelet/glamour"
)

func Markdown(html string, width int, title string) string {
	opt := &md.Options{
		EscapeMode: "disabled",
	}
	converter := md.NewConverter("", true, opt)
	converter.Use(plugin.YoutubeEmbed(), plugin.GitHubFlavored())
	body, err := converter.ConvertString(html)
	if err != nil {
		log.Fatal(err)
	}

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(width),
		mdStyle,
	)
	if err != nil {
		log.Fatal(err)
	}

	if len(title) > 0 {
		body = fmt.Sprintf("# %s\n---\n%s", title, body)
	}

	out, err := renderer.Render(body)
	if err != nil {
		log.Fatal(err)
	}
	return out
}

var mdStyle = glamour.WithStylesFromJSONBytes(
	[]byte(`
        {
            "document": {
                "margin": 0
            }
        }
    `),
)
