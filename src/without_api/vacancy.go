package withoutapi

import (
	"fmt"

	"github.com/gocolly/colly"
)



var headers = map[string]string {
	"authority": "hh.ru",
	"accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
	"accept-language": "ru,en;q=0.9,en-GB;q=0.8,en-US;q=0.7",
	"cache-control": "max-age=0",
	"cookie": "__ddg1_=rZzOnXzyl1w4Iqeoi6Pg; _xsrf=f56752d0315b1eb1b989e5f3a684037c; regions=1; crypted_hhuid=03AD38DFCA4BE31F7559144CA4E70C4BCC27F22B084F11A888E78D41F260DEF6; hhtoken=QI0gZ5Ftopb5QxOdQdMZWWa4qsiz; hhuid=M9g!7T2R3Y2m3WP0k1Ezaw--; GMT=3; tmr_lvid=69cc1ec94b709f774d578e2102edad40; tmr_lvidTS=1676972881719; _ym_uid=1676972882198158643; _ym_d=1676972882; _ym_isad=1; _gid=GA1.2.707814978.1676973060; iap.uid=2040634ba3a449488728851c1c02d995; region_clarified=NOT_SET; display=desktop; _ym_visorc=w; hhrole=applicant; crypted_id=C7DD4A3472D3E4DA3D840DCD927194E1A5471C6F722104504734B32F0D03E9F8; hhul=f4fb6bf75411f24a451706336e8411926199e4529664bbfd0bfc2b8342117bfa; total_searches=2; tmr_detect=1%7C1676978505426; device_breakpoint=l; _hi=148391044; _ga=GA1.2.2022095896.1676972882; __zzatgib-w-hh=MDA0dC0jViV+FmELHw4/aQsbSl1pCENQGC9LX3NbOx8kYUxdUHUTTX5bSxh/J1gNDQ5jREVzdS88ISQbOVURCxIXRF5cVWl1FRpLSiVueCplJS0xViR8SylEW1R/LiAZfGsrVw4NVy8NPjteLW8PKhMjZHYhP04hC00+KlwVNk0mbjN3RhsJHlksfEspNVhQMylNFndxK1IPFBZzcnVzKXFmIBxNWlFMVQl5LU0ZeGsjUgwLYT8zaWVpcC9gIBIlEU1HGEVkW0I2KBVLcU8cenZffSpBbCNoTmEkQ11VfycVe0M8YwxxFU11cjgzGxBhDyMOGFgJDA0yaFF7CT4VHThHKHIzd2UqP20gaEhZI0dHSWtlTlNCLGYbcRVNCA00PVpyIg9bOSVYCBI/CyYfGH10KlcMC2NFR3RvG382XRw5YxEOIRdGWF17TEA=MooZOQ==; _ga_44H5WGZ123=GS1.1.f42d2dae01ae01415e94e4406989eccad8f7652471520fbd7cd607f82427b106.2.1.1676978512.19.0.0; cfidsgib-w-hh=qzhULnCqxxybLwLuHNI///1lvjkmV2riLKaZrNaPsk0m7s/MBWpXrniYCLZO7qRAGv4s7r40UHaVgm99AA1NxdY9q3+n6y241j+2HSTRp6Z2PFoynswPdKAvWP/uXCO3I4gaHHnvJLXlrtWG8xn0PbG3hZVLx04VqbdktQ==; cfidsgib-w-hh=qzhULnCqxxybLwLuHNI///1lvjkmV2riLKaZrNaPsk0m7s/MBWpXrniYCLZO7qRAGv4s7r40UHaVgm99AA1NxdY9q3+n6y241j+2HSTRp6Z2PFoynswPdKAvWP/uXCO3I4gaHHnvJLXlrtWG8xn0PbG3hZVLx04VqbdktQ==; gsscgib-w-hh=isBjK0ju2J5Gy5Bmtr6GzWq0zKUvpqlbzMWgMjrzx9Afi+4POS33tP3j4dkNAj+JIs+ZQMsyX4gPpYgr8rLmn24UFhC3dMy5UHRaQHx4UT0QueKXOTRTQvjcKQqFqAERaBvc2so+qocOwVfuXURgeXkOqiElDzSrmq/OTEcB3/8ACJKYM7pD2EJrnrXV12upVRNVdBiCie5Qis2JHasGyeH7lfEB2c9o3TFDhNKVkibUWPfM6j6De45iOppdo7vDllE=; fgsscgib-w-hh=22Pn9a9642052e2837efe8813a11e5d9ccbcd97b",
	"referer": "https://hh.ru/vacancies/naladchikom_kipia",
	"sec-ch-ua": `"Chromium";v="110", "Not A(Brand";v="24", "Microsoft Edge";v="110"`,
	"sec-ch-ua-mobile": "?0",
	"sec-ch-ua-platform": `"Windows"`,
	"sec-fetch-dest": "document",
	"sec-fetch-mode": "navigate",
	"sec-fetch-site": "same-origin",
	"sec-fetch-user": "?1",
	"upgrade-insecure-requests": "1",
	"user-agent": "User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.217 1.95 Safari/537.36",

}

func GetVacancy(url string) {
	body := getBody(url)
	fmt.Println(body.ChildText(`h1[data-qa=vacancy-title]`))

}



func getBody(url string) (body *colly.HTMLElement) {
	c := colly.NewCollector()
	c.OnHTML("div.main-content", func(h *colly.HTMLElement) {
		body = h
	})
	err := c.Post(url, headers)
	checkErr(err)
	return
}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}