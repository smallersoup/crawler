package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
	"log"
)

var (
	// <td><span class="label">年龄：</span>25岁</td>
	ageReg = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
	// <td><span class="label">身高：</span>182CM</td>
	heightReg = regexp.MustCompile(`<td><span class="label">身高：</span>(\d+)CM</td>`)
	// <td><span class="label">月收入：</span>5001-8000元</td>
	incomeReg = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
	//<td><span class="label">婚况：</span>未婚</td>
	marriageReg = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
	//<td><span class="label">学历：</span>大学本科</td>
	educationReg = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
	//<td><span class="label">工作地：</span>安徽蚌埠</td>
	workLocationReg = regexp.MustCompile(`<td><span class="label">工作地：</span>([^<]+)</td>`)
	// <td><span class="label">职业：</span>--</td>
	occupationReg = regexp.MustCompile(`<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`)
	//  <td><span class="label">星座：</span>射手座</td>
	xinzuoReg = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
	//<td><span class="label">籍贯：</span>上海徐汇</td>
	hokouReg = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
	// <td><span class="label">住房条件：</span><span field="">--</span></td>
	houseReg = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
	// <td><span class="label">性别：</span><span field="">([^<]+)</span></td>
	genderReg = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)

	// <td><span class="label">体重：</span><span field="">67KG</span></td>
	weightReg = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
	//<h1 class="ceiling-name ib fl fs24 lh32 blue">怎么会迷上你</h1>
	//nameReg = regexp.MustCompile(`<h1 class="ceiling-name ib fl fs24 lh32 blue">([^\d]+)</h1>  `)
	//<td><span class="label">是否购车：</span><span field="">未购车</span></td>
	carReg = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

	guessRe = regexp.MustCompile(`href="(http://album.zhenai.com/u/\w+)"[^>]*>([^<]+)</a>`)
	idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)
)

func parseProfile(contents []byte, url, name string) engine.ParserResult {

	profile := model.Profile{}

	age, err := strconv.Atoi(extractString(contents, ageReg))

	if err != nil {
		profile.Age = 0
	} else {
		profile.Age = age
	}

	height, err := strconv.Atoi(extractString(contents, heightReg))
	if err != nil {
		profile.Height = 0
	} else {
		profile.Height = height
	}

	weight, err := strconv.Atoi(extractString(contents, weightReg))
	if err != nil {
		profile.Weight = 0
	} else {
		profile.Weight = weight
	}

	profile.Income = extractString(contents, incomeReg)

	profile.Car = extractString(contents, carReg)

	profile.Education = extractString(contents, educationReg)
	profile.Gender = extractString(contents, genderReg)

	profile.Hukou = extractString(contents, hokouReg)
	profile.Income = extractString(contents, incomeReg)
	profile.Marriage = extractString(contents, marriageReg)
	profile.Name = name
	profile.Occupation = extractString(contents, occupationReg)
	profile.WorkLocation = extractString(contents, workLocationReg)
	profile.Xinzuo = extractString(contents, xinzuoReg)
	profile.House = extractString(contents, houseReg)

	result := engine.ParserResult{
		Items: []engine.Item{
			{
				Type:    "zhenai",
				Id:      extractString([]byte(url), idUrlRe),
				Url:     url,
				Payload: profile,
			},
		},
	}

	matches := guessRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		//url := string(m[1])
		//name := string(m[2])

		log.Printf("guess url:%s name:%s\n", url, name)
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			//Parser: NewProfileParser(string(m[2])),
			ParserFunc: ProfileParser(string(m[2])),
		})
	}

	return result
}

//get value by reg from contents
func extractString(contents []byte, re *regexp.Regexp) string {

	m := re.FindSubmatch(contents)

	if len(m) > 0 {
		return string(m[1])
	} else {
		return ""
	}
}
