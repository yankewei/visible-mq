package models

type TopicModel struct {
}

type TopicData struct {
	Id int
	Topic string
	Class string
	Function string
}

func (Topic *TopicModel) GetAll() {
	rows, err := DB.Query("SELECT * FROM " + Topic.GetTable())
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {

	}
}

func (Topic *TopicModel) GetTable() string {
	return "topic"
}

func (Topic *TopicModel) Insert(data *TopicData) int64 {
	stmt, err := DB.Prepare("INSERT INTO " + Topic.GetTable() + "VALUES(?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(data.Topic, data.Class, data.Function)
	if err != nil {
		panic(err)
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	return lastId
}

func (topic *TopicModel) GetListByChannel(channel string) []TopicData {
	topicSubscribeList := []TopicData{}
	rows, _ := DB.Query("SELECT * FROM " + topic.GetTable() + " WHERE topic = ?", channel)
	defer rows.Close()
	for rows.Next() {
		var tmp TopicData
		rows.Scan(&tmp.Id, &tmp.Topic, &tmp.Class, &tmp.Function)
		topicSubscribeList = append(topicSubscribeList, tmp)
	}
	if len(topicSubscribeList) != 0 {
		return topicSubscribeList
	}
	return []TopicData{}
}