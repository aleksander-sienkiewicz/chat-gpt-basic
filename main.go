package main

//this proj, sends a prompt to hpt and outputs it.
import ( //were gonna nneed a bunch of imports
	"context"
	"fmt"
	"log"
	"os"

	"github.com/PullRequestInc/go-gpt3" //this is a gpt client - has a lot of active dev rn and v popular
	"github.com/joho/godotenv"          //for env file
	//when importing use go mod tidy to remove these import squigglies :D
)

func main() {
	godotenv.Load()

	apiKey := os.Getenv("API_KEY") //get env func, api key is gitIgnored from env file
	if apiKey == "" {              //if theres no key in the var from env
		log.Fatalln("Missing API KEY") //uses log import
	}

	ctx := context.Background()      ///context import: https://www.youtube.com/watch?time_continue=45&v=h2RdcrMLQAo&embeds_euri=https%3A%2F%2Fwww.google.com%2Fsearch%3Fq%3Dwhat%2Bis%2Bcontext%2Bgolang%26sxsrf%3DAPwXEdeAT_HQ2Cw48QLbzc5bl1wpJgxhEA%3A1682804671580%26source%3Dlnms%26tbm%3Dvi&source_ve_path=MjM4NTE&feature=emb_title
	client := gpt3.NewClient(apiKey) //

	resp, err := client.Completion(ctx, gpt3.CompletionRequest{ //hit completions aps
		//required for completion,
		Prompt:    []string{"The first thing you should know about javascript is"},
		MaxTokens: gpt3.IntPtr(30),
		Stop:      []string{"."},
		Echo:      true,
	})
	if err != nil { //handle error
		log.Fatalln(err) //log fatal
	}
	fmt.Println(resp.Choices[0].Text) //if all goes well, print out response
}
