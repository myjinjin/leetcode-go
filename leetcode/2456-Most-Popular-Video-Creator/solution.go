package leetcode

type viewInfo struct {
	minID     string
	maxView   int
	totalView int
}

func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
	creatorViews := make(map[string]*viewInfo)
	maxView := 0
	for i, creator := range creators {
		if v, ok := creatorViews[creator]; ok {
			v.totalView += views[i]
			if views[i] > v.maxView || (views[i] == v.maxView && ids[i] < v.minID) {
				v.maxView = views[i]
				v.minID = ids[i]
			}
		} else {
			creatorViews[creator] = &viewInfo{
				minID:     ids[i],
				maxView:   views[i],
				totalView: views[i],
			}
		}
		maxView = max(maxView, creatorViews[creator].totalView)
	}

	result := [][]string{}

	for k, v := range creatorViews {
		if v.totalView == maxView {
			result = append(result, []string{k, v.minID})
		}
	}

	return result
}
