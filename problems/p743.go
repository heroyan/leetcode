package problems

//https://leetcode-cn.com/problems/network-delay-time/
func networkDelayTime(times [][]int, n int, k int) int {
	maxInt := int(^uint(0) >> 1)
	weight := map[int]map[int]int{}
	adjs := map[int][]int{}
	// 构造图的权重和邻接表
	for _, cols := range times {
		if _, ok := weight[cols[0]]; !ok {
			weight[cols[0]] = map[int]int{}
		}
		weight[cols[0]][cols[1]] = cols[2]
		if _, ok := adjs[cols[0]]; !ok {
			adjs[cols[0]] = []int{}
		}
		adjs[cols[0]] = append(adjs[cols[0]], cols[1])
	}

	var distTo []int
	for i := 0; i < n+1; i++ {
		distTo = append(distTo, maxInt)
	}
	distTo[k] = 0
	// TODO：这里可以做一个优化，使用优先级队列，减少下面node append到queue的次数
	var queue []State
	queue = append(queue, State{id: k, distFromStart: 0})
	for len(queue) > 0 {
		curNode := queue[0].id
		curDist := queue[0].distFromStart
		// 相当于出队
		queue = queue[1:]
		if curDist > distTo[curNode] {
			// 已经有更短的路径到达curNode了，这条路径忽略，这里只是加速算法，对结果其实没有影响
			// 不continue也可以
			continue
		}
		for _, nextNode := range adjs[curNode] {
			distNext := distTo[curNode] + weight[curNode][nextNode]
			if distNext < distTo[nextNode] {
				distTo[nextNode] = distNext
				queue = append(queue, State{id: nextNode, distFromStart: distNext})
			}
		}
	}

	ret := -1
	for idx, m := range distTo {
		if idx == 0 {
			continue
		}
		if m == maxInt {
			// 有些点不可达，则直接返回-1了
			return -1
		}
		if m > ret {
			ret = m
		}
	}

	return ret
}

type State struct {
	id int
	distFromStart int
}

// 最近真是诸事不顺，一家人轮着生病，好了一轮又一轮。成年人的生活没有容易二字，深有体会。
// 自己生个病还好，特别是小孩生病，比自己生病还难受。他的一举一动都牵动着你，大声的咳嗽、流鼻涕、拉肚子、发烧，让你手足无措欲哭无泪。
// 有时候真想狠狠的扇自己两个巴掌，我觉得小孩的生病都是因为家长没有带好，比如没穿好衣服、没盖好被子、给小孩吃坏东西了。上周末独自一人带着
// 豆豆去公园玩了一趟，甚至我以为就是我把他弄感冒了，风大了吹的，小孩玩水湿了裤子冻的，种种可能情况涌上心头。