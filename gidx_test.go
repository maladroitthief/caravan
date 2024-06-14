package caravan_test

import (
	"math/rand"
	"testing"

	"github.com/maladroitthief/caravan"
)

const (
	ContainerSize = 1000000
)

func BenchmarkGenerationalIndexIndexing(b *testing.B) {
	type entity struct {
		id int
	}
	genIndexAllocator := *caravan.NewGIDXAllocator()
	genIndexArray := caravan.NewGIDXArray[entity]()
	indexIds := []caravan.GIDX{}
	for i := 0; i < ContainerSize; i++ {
		entityId := genIndexAllocator.Allocate()
		indexIds = append(indexIds, entityId)
		genIndexArray = genIndexArray.Set(entityId, entity{id: rand.Int()})
	}

	for n := 0; n < b.N; n++ {
		_ = genIndexArray.Get(indexIds[rand.Intn(ContainerSize)])
	}
}

func BenchmarkGenerationalIndexInsert(b *testing.B) {
	type entity struct {
		id int
	}
	genIndexAllocator := *caravan.NewGIDXAllocator()
	genIndexArray := caravan.NewGIDXArray[entity]()
	var entityId caravan.GIDX

	for n := 0; n < b.N; n++ {
		entityId = genIndexAllocator.Allocate()
		genIndexArray = genIndexArray.Set(entityId, entity{id: rand.Int()})
	}
}

// func BenchmarkGenerationalIndexDelete(b *testing.B) {
// 	type entity struct {
// 		id int
// 	}
// 	genIndexAllocator := *caravan.NewGIDXAllocator()
// 	genIndexArray := caravan.NewGIDXArray[entity]()
// 	indexIds := []caravan.GIDX{}
// 	for i := 0; i < ContainerSize; i++ {
// 		entityId := genIndexAllocator.Allocate()
// 		indexIds = append(indexIds, entityId)
// 		genIndexArray = genIndexArray.Set(entityId, entity{id: rand.Int()})
// 	}

// 	for n := 0; n < b.N; n++ {
// 		if n < ContainerSize {
// 			index := rand.Intn(len(indexIds))

// 			genIndexArray = genIndexArray.Remove(indexIds[index])
// 			ge.Remove(indexIds[index])

// 			indexIds[index] = indexIds[len(indexIds)-1]
// 			indexIds = indexIds[:len(indexIds)-1]
// 		}
// 	}
// }

func BenchmarkGIDXPoolIndexing(b *testing.B) {
	type entity struct {
		id int
	}
	pool := *caravan.NewGIDXPool[entity]()
	indexIds := []caravan.GIDX{}
	for i := 0; i < ContainerSize; i++ {
		id := pool.Allocate()
		indexIds = append(indexIds, id)
		pool.Set(id, entity{id: rand.Int()})
	}

	for n := 0; n < b.N; n++ {
		_ = pool.Get(indexIds[rand.Intn(ContainerSize)])
	}
}

func BenchmarkGIDXInsert(b *testing.B) {
	type entity struct {
		id int
	}

	pool := *caravan.NewGIDXPool[entity]()
	var id caravan.GIDX
	for n := 0; n < b.N; n++ {
		id = pool.Allocate()
		pool.Set(id, entity{id: rand.Int()})
	}
}

func BenchmarkGIDXPoolDelete(b *testing.B) {
	type entity struct {
		id int
	}

	pool := *caravan.NewGIDXPool[entity]()
	indexIds := []caravan.GIDX{}
	for i := 0; i < ContainerSize; i++ {
		id := pool.Allocate()
		indexIds = append(indexIds, id)
		pool.Set(id, entity{id: rand.Int()})
	}

	for n := 0; n < b.N; n++ {
		if n < ContainerSize {
			index := rand.Intn(len(indexIds))

			pool.Remove(indexIds[index])

			indexIds[index] = indexIds[len(indexIds)-1]
			indexIds = indexIds[:len(indexIds)-1]
		}
	}
}
