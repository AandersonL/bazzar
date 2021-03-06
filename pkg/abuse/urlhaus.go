package abuse

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aandersonl/bazzar/pkg/utils"
)

func QueryLast(limit int) (string, LastUrls) {
	recentUrl := URLHAUS_API_URL + "s/recent/limit/" + fmt.Sprint((limit)) + "/"

	resp, err := http.Get(recentUrl)
	utils.ExitIfError(err)

	body, err := ioutil.ReadAll(resp.Body)
	utils.ExitIfError(err)

	last := LastUrls{}

	err = json.Unmarshal(body, &last)
	utils.ExitIfError(err)

	return string(body), last
}

func QueryUrl(url string) (string, URLResponse) {
	query_url_post.Set("url", url)

	resp, err := http.PostForm(URLHAUS_API_URL, query_url_post)
	utils.ExitIfError(err)

	body, err := ioutil.ReadAll(resp.Body)
	utils.ExitIfError(err)

	urlResponse := URLResponse{}

	err = json.Unmarshal(body, &urlResponse)
	utils.ExitIfError(err)

	return string(body), urlResponse
}

func QueryHost(host string) (string, HostResponse) {
	query_host_post.Set("host", host)

	resp, err := http.PostForm(URLHAUS_API_HOST, query_host_post)
	utils.ExitIfError(err)

	body, err := ioutil.ReadAll(resp.Body)
	utils.ExitIfError(err)

	hostResponse := HostResponse{}
	err = json.Unmarshal(body, &hostResponse)
	utils.ExitIfError(err)

	return string(body), hostResponse

}
