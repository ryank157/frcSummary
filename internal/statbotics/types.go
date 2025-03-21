package statbotics

import "net/http"

// Client is a struct that holds the configuration for the Statbotics API client.
type Client struct {
	baseURL    string
	httpClient *http.Client
}
type DefaultResponse struct {
	Name string `json:"name"`
}

// MatchResponse defines the structure of the response for the match route.
type MatchResponse struct {
	Key           string      `json:"key"`
	Year          int         `json:"year"`
	Event         string      `json:"event"`
	Week          int         `json:"week"`
	Elim          bool        `json:"elim"`
	CompLevel     string      `json:"comp_level"`
	SetNumber     int         `json:"set_number"`
	MatchNumber   int         `json:"match_number"`
	MatchName     string      `json:"match_name"`
	Time          int64       `json:"time"`
	PredictedTime int64       `json:"predicted_time"`
	Status        string      `json:"status"`
	Video         interface{} `json:"video"` // This is confirmed to be null or potentially a string, leave as interface{}

	Alliances struct {
		Red struct {
			TeamKeys          []int         `json:"team_keys"`           // Confirmed to be integers
			SurrogateTeamKeys []interface{} `json:"surrogate_team_keys"` // Empty Arrays
			DqTeamKeys        []interface{} `json:"dq_team_keys"`        // Empty Arrays
		} `json:"red"`
		Blue struct {
			TeamKeys          []int         `json:"team_keys"`           // Confirmed to be integers
			SurrogateTeamKeys []interface{} `json:"surrogate_team_keys"` // Empty Arrays
			DqTeamKeys        []interface{} `json:"dq_team_keys"`        // Empty Arrays

		} `json:"blue"`
	} `json:"alliances"`
	Pred struct {
		Winner      string  `json:"winner"`
		RedWinProb  float64 `json:"red_win_prob"`
		RedScore    float64 `json:"red_score"`
		BlueScore   float64 `json:"blue_score"`
		RedAutoRp   float64 `json:"red_auto_rp"`
		BlueAutoRp  float64 `json:"blue_auto_rp"`
		RedCoralRp  float64 `json:"red_coral_rp"`
		BlueCoralRp float64 `json:"blue_coral_rp"`
		RedBargeRp  float64 `json:"red_barge_rp"`
		BlueBargeRp float64 `json:"blue_barge_rp"`
		RedRp1      float64 `json:"red_rp_1"`
		BlueRp1     float64 `json:"blue_rp_1"`
		RedRp2      float64 `json:"red_rp_2"`
		BlueRp2     float64 `json:"blue_rp_2"`
		RedRp3      float64 `json:"red_rp_3"`
		BlueRp3     float64 `json:"blue_rp_3"`
	} `json:"pred"`
	Result struct {
		Winner                   *string     `json:"winner"` // Can be string or null
		RedScore                 int         `json:"red_score"`
		BlueScore                int         `json:"blue_score"`
		RedNoFoul                interface{} `json:"red_no_foul"`         // Could be int or null
		BlueNoFoul               interface{} `json:"blue_no_foul"`        // Could be int or null
		RedAutoPoints            interface{} `json:"red_auto_points"`     // Could be int or null
		BlueAutoPoints           interface{} `json:"blue_auto_points"`    // Could be int or null
		RedTeleopPoints          interface{} `json:"red_teleop_points"`   // Could be int or null
		BlueTeleopPoints         interface{} `json:"blue_teleop_points"`  // Could be int or null
		RedEndgamePoints         interface{} `json:"red_endgame_points"`  // Could be int or null
		BlueEndgamePoints        interface{} `json:"blue_endgame_points"` // Could be int or null
		RedAutoRp                bool        `json:"red_auto_rp"`
		BlueAutoRp               bool        `json:"blue_auto_rp"`
		RedCoralRp               bool        `json:"red_coral_rp"`
		BlueCoralRp              bool        `json:"blue_coral_rp"`
		RedBargeRp               bool        `json:"red_barge_rp"`
		BlueBargeRp              bool        `json:"blue_barge_rp"`
		RedTiebreakerPoints      int         `json:"red_tiebreaker_points"`
		BlueTiebreakerPoints     int         `json:"blue_tiebreaker_points"`
		RedAutoLeavePoints       interface{} `json:"red_auto_leave_points"`  // Could be int or null
		BlueAutoLeavePoints      interface{} `json:"blue_auto_leave_points"` // Could be int or null
		RedAutoCoral             interface{} `json:"red_auto_coral"`         // Could be int or null
		BlueAutoCoral            interface{} `json:"blue_auto_coral"`        // Could be int or null
		RedAutoCoralPoints       interface{} `json:"red_auto_coral_points"`
		BlueAutoCoralPoints      interface{} `json:"blue_auto_coral_points"`
		RedTeleopCoral           interface{} `json:"red_teleop_coral"`
		BlueTeleopCoral          interface{} `json:"blue_teleop_coral"`
		RedTeleopCoralPoints     interface{} `json:"red_teleop_coral_points"`
		BlueTeleopCoralPoints    interface{} `json:"blue_teleop_coral_points"`
		RedCoralL1               interface{} `json:"red_coral_l1"`
		BlueCoralL1              interface{} `json:"blue_coral_l1"`
		RedCoralL2               interface{} `json:"red_coral_l2"`
		BlueCoralL2              interface{} `json:"blue_coral_l2"`
		RedCoralL3               interface{} `json:"red_coral_l3"`
		BlueCoralL3              interface{} `json:"blue_coral_l3"`
		RedCoralL4               interface{} `json:"red_coral_l4"`
		BlueCoralL4              interface{} `json:"blue_coral_l4"`
		RedTotalCoralPoints      interface{} `json:"red_total_coral_points"`
		BlueTotalCoralPoints     interface{} `json:"blue_total_coral_points"`
		RedProcessorAlgae        interface{} `json:"red_processor_algae"`
		BlueProcessorAlgae       interface{} `json:"blue_processor_algae"`
		RedProcessorAlgaePoints  interface{} `json:"red_processor_algae_points"`
		BlueProcessorAlgaePoints interface{} `json:"blue_processor_algae_points"`
		RedNetAlgae              interface{} `json:"red_net_algae"`
		BlueNetAlgae             interface{} `json:"blue_net_algae"`
		RedNetAlgaePoints        interface{} `json:"red_net_algae_points"`
		BlueNetAlgaePoints       interface{} `json:"blue_net_algae_points"`
		RedTotalAlgaePoints      interface{} `json:"red_total_algae_points"`
		BlueTotalAlgaePoints     interface{} `json:"blue_total_algae_points"`
		RedTotalGamePieces       interface{} `json:"red_total_game_pieces"`
		BlueTotalGamePieces      interface{} `json:"blue_total_game_pieces"`
		RedBargePoints           interface{} `json:"red_barge_points"`  // Could be int or null
		BlueBargePoints          interface{} `json:"blue_barge_points"` // Could be int or null
		RedRp1                   bool        `json:"red_rp_1"`
		BlueRp1                  bool        `json:"blue_rp_1"`
		RedRp2                   bool        `json:"red_rp_2"`
		BlueRp2                  bool        `json:"blue_rp_2"`
		RedRp3                   bool        `json:"red_rp_3"`
		BlueRp3                  bool        `json:"blue_rp_3"`
	} `json:"result"`
}
