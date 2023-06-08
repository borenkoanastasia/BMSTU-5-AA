package main
import "math"
import "math/rand"
import "time"
//import "fmt"


const START_PHEROMON = 0.5
const EPS = 0;//0.00005;

func mremove(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}
type ant_t struct{
	J []int //cities_to_go
	end int
	//ans []int
	pos int
}

type colony_t struct{
	D [][]float64 
	Tao [][]float64 
	alpha float64 
	betta float64 
	q float64 
	p float64 
	t_max int
}

func (e *colony_t)generate_colony(D [][]float64, alpha float64, betta float64, q float64, p float64, t_max int) {
	e.D = D 
	e.alpha = alpha 
	e.betta = betta 
	e.p = p 
	e.q = q 
	e.t_max = t_max
	e.generate_pheromon()
}
func (e *colony_t)generate_pheromon() {
	e.Tao = make_float_matrix(len(e.D), len(e.D[0]))
	for i:=0;i<len(e.D);i++{
		for j:=0;j<len(e.D[i]);j++{
			e.Tao[i][j] = START_PHEROMON
		}
	}
}
func (e colony_t)generate_ant_in_city(city_i int)ant_t{
	var ant ant_t

	ant.J = make([]int, 0)
	ant.J = make([]int, 0)
	for i:=0;i<len(e.D);i++{
		if (i == city_i){
			continue
		}
		ant.J = append(ant.J, i)
	}
	ant.pos = city_i
	ant.end = ant.pos
	return ant
}
func (e colony_t)generate_all_ants()[]ant_t{
	var ants []ant_t = make([]ant_t, len(e.D))
	for i:=0;i<len(e.D);i++{
		ants[i] = e.generate_ant_in_city(i)
	}
	return ants
}

func get_spec_summ(e colony_t, a ant_t)float64{
	var sum float64 = 0
	for l:=0;l<len(a.J);l++{
		if (e.D[a.pos][a.J[l]]<=0){
			sum +=0
			continue
		}
		var n_ij = 1/e.D[a.pos][a.J[l]]
		var pheromon = e.Tao[a.pos][a.J[l]]
		sum += math.Pow(pheromon, e.alpha) * math.Pow(n_ij, e.betta)
	}
	return sum
}

func get_route(e colony_t, a *ant_t, l_k *float64)int{
	var p_ijk = make([]float64, len(a.J))
	var chis float64
	chis = 1
	var sum float64 = get_spec_summ(e, *a)
	if (math.Abs(sum) < EPS){
		return -1
	}
	
	for i:=0;i<len(a.J);i++{
		if (math.Abs(e.D[a.pos][a.J[i]])<=EPS){
			if (i > 0){
				p_ijk[i] = p_ijk[i-1]
			} else{
				p_ijk[i] = 0
			}
			continue
		}
		var n_ij = 1/e.D[a.pos][a.J[i]]
		var pheromon = e.Tao[a.pos][a.J[i]]
		chis = math.Pow(pheromon, e.alpha) * math.Pow(n_ij, e.betta)
		if (i > 0){
			p_ijk[i] = p_ijk[i-1] + chis/sum*100
		} else{
		    p_ijk[i] = chis/sum*100
		}
	}
	var rand_num = rand.Intn(100)

	for i:=0;i<len(p_ijk);i++{
		if (float64(rand_num)  <= p_ijk[i]){
			var result = a.J[i]
			a.J = mremove(a.J, i)
			*l_k = e.D[a.pos][result]
			a.pos = result

			return result
		}
	}
	return -1
}

func (e *colony_t)update_pheromons(a []ant_t, l[]float64) {
	var del_t float64
	del_t = 0
	for k:=0;k<len(e.Tao);k++{
		for i:=0;i< len(e.Tao[k]);i++{
			if e.D[k][i] != 0{
				if l[k] >0{
					del_t = e.q / float64(e.D[k][i])
				} else {
					del_t = 0
				}
				e.Tao[k][i] = (1-e.p) * (float64(1)) + del_t
			}
			if e.Tao[k][i] <= 0{
				e.Tao[k][i] = 0.1
			}
		}
	}
	
}

func ant_alg(e colony_t, l_res *float64)[]int{
	rand.Seed(time.Now().UnixNano())
	var ants []ant_t
	var l_k_res []float64 = make([]float64, len(e.D))
	var t_res []int 
	*l_res = -1
	for i:=0;i<e.t_max;i++{
		ants = e.generate_all_ants() 
		for j:=0;j<len(ants);j++{
			var l_k float64 = 0
			var t_k[]int
			t_k = append(t_k, ants[j].pos)
			var succes_route = true
			for l:=0;len(ants[j].J)>0; l++{
				var l_k_cur float64
				t_k = append(t_k,get_route(e, &(ants[j]), &l_k_cur))
				if (t_k[len(t_k)-1] == -1){
					//fmt.Println("BAD ERROR, cant go")
					succes_route = false
					l_k_res[j] = -1
					break
				}
				l_k += l_k_cur  
			}
			if (succes_route){
				if (e.D[t_k[0]][t_k[len(t_k)-1]] <=0){
					//fmt.Println("BAD ERROR, cant go")
					succes_route = false
					l_k_res[j] = -1
					continue
				}
				l_k += e.D[t_k[0]][t_k[len(t_k)-1]]
				l_k_res[j] = l_k
				if (*l_res > l_k || *l_res == -1){
					*l_res = l_k 
					t_res = t_k
			    }
			}
			
		}
		e.update_pheromons(ants, l_k_res)
	}
	return t_res
}

func AllAntsAlg(m float_matrix_t, l *float64)[]int{
	var colony colony_t;
	colony.generate_colony(m, 3.0, 7.0, 20.0, 0.6, 10)

	var ans =ant_alg(colony, l)
	return ans
}


/*
type colony_t struct{
	D [][]float64 // собственно весы
	Tao [][]float64 // феромоны
	alpha float64 //приоритет пути
	betta float64 //приоритет феромона
	q float64 // переносимый муравьем феромона
	p float64 // коэффициент испарения феромона
	t_max int // максимальное время жизни колонии
}*/

/*
type env_t struct{
	graph float_matrix_t
	pheromon float_matrix_t
	alpha float64 //приоритет пути
	betta float64 //приоритет феромона
	q float64 // переносимый муравьем феромона
	p float64 // коэффициент испарения феромона
	days_count int
}

const START_PHEROMON = 0.5

func (e env_t)generate_pheromon() {
	e.pheromon = make_float_matrix(len(e.graph), len(e.graph[0]))
	for i:=0;i<len(e.graph);i++{
		for j:=0;j<len(e.graph[i]);j++{
			e.pheromon[i][j] = START_PHEROMON
		}
	}
}
func make_env(graph float_matrix_t, alpha float64, betta float64, q float64, p float64, days_count int)*env_t{
	var e = new(env_t)
	e.graph = graph 
	e.generate_pheromon()
	e.alpha = alpha 
	e.betta = betta
	e.p = p
	e.q = q
	e.days_count = days_count
	return e
}

func generate_ants(e env_t){
	var cities_count = len(e.graph)*len(e.graph[0])
	for i:=0;i<cities_count;i++{
		make_new_ant()
		update_pheromon()
	}
}
func find_solves(){}
func update_pheromon(){}


func main_myravs(e env_t){

	//initialise_pheromon() случается при генерации колонии
	for j:=0;j<e.days_count;j++{
		generate_ants(e)
		find_solves()
		update_pheromon()
	}
}



/*
//import "math/rand"
//import "time"
//import "math"

// Объявляем важные константы и типы
const PHEROMON_MIN = 0.5

type ant_t struct{
	way int_array_t   //не пройденные вершины
	route int_array_t //пройденные вершины
}

type ant_colony_t struct{
	graph float_matrix_t
	pheromonths float_matrix_t
	alpha float64
	beta float64
	Q float64
	p float64
}

func calculate_Q(graph float_matrix_t)float64{
	var q float64 = 0
	for i:=0;i<len(graph);i++{
		for j:=0;j<len(graph);j++{
			q += graph[i][j]
		}
	}
	return q
}

func (acolony ant_colony_t)make_feramons(graph float_matrix_t, p float64, alpha float64, beta float64){
	var m = make_float_matrix(len(graph), len(graph))
	for i:=0;i<len(graph);i++{
		for j:=0;j<len(graph);j++{
			m[i][j] = PHEROMON_MIN
		}
	}
	acolony.pheromonths = m
	acolony.graph = graph
	acolony.p = p
	acolony.alpha = alpha
	acolony.beta = beta 
}

func myravs(graph float_matrix_t, day_count int, p float64, alpha float64, beta float64){}

/*
const PHEROMON_MIN = 0.5

func (acolony ant_colony_t)make_feramons(size int, p float64, ){
	var m matrix_t = make_empty_matrix(size, size)
	for i:=0;i<size;i++{
		for j:=0;j<size;j++{
			m.elem[i][j] = PHEROMON_MIN
		}
	}
	acolony.pheromonths = m
}

func myravs(){}
*/

