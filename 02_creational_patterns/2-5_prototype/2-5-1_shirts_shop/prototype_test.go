package prototype

import "testing"

func TestClone(t *testing.T) {
	shirtCache := GetShirtsCloner()
	if shirtCache == nil {
		t.Fatal("Received cache was nil")
	}

	whiteItem1, err := shirtCache.GetClone(White)
	if err != nil {
		t.Error(err)
	}

	if whiteItem1 == whitePrototype {
		t.Error("item1 cannot be equal to the white prototype")
	}

	whiteShirt1, ok := whiteItem1.(*Shirt)
	if !ok {
		t.Fatal("Type assertion for shirt1 couldn't be done successfully")
	}

	if whiteShirt1.Color != White {
		t.Errorf("Expected shirt color is White, but found %d", whiteShirt1.Color)
	}

	whiteItem2, err := shirtCache.GetClone(White)
	if err != nil {
		t.Error(err)
	}

	whiteShirt2, ok := whiteItem2.(*Shirt)
	if !ok {
		t.Fatal("Type assertion for shirt2 couldn't be done successfully")
	}

	whiteShirt1.SKU = "abbcc"
	if whiteShirt1.SKU == whiteShirt2.SKU {
		t.Error("SKU's of shirt1 and shirt2 must be different")
	}

	if whiteShirt1 == whiteShirt2 {
		t.Error("Shirt1 cannot be equal to shirt2")
	}

	t.Logf("LOG: %s", whiteShirt1.GetInfo())
	t.Logf("LOG: %s", whiteShirt2.GetInfo())

	t.Logf("LOG: The memory positions of the shirts are different %p != %p\n\n", &whiteShirt1, &whiteShirt2)
}
