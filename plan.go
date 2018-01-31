package goap

type Plan []Action

func (p *Plan) Cost() int {
	cost := 0
	for _, action := range *p {
		cost += action.Cost()
	}
	return cost
}
