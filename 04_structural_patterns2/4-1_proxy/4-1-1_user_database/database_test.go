package database

import (
	"math/rand"
	"testing"
)

func Test_UserListProxy(t *testing.T) {
	someDatabase := UserList{}

	for i := 0; i < 1000000; i++ {
		n := rand.Int31()
		someDatabase = append(someDatabase, User{ID: n})
	}

	proxy := UserListProxy{
		SomeDatabase:  &someDatabase,
		StackCache:    UserList{},
		StackCapacity: 2,
	}

	knownIDs := [3]int32{
		someDatabase[3].ID,
		someDatabase[4].ID,
		someDatabase[5].ID,
	}

	t.Run("FindUser - Empty Cache", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}

		if user.ID != knownIDs[0] {
			t.Error("Returned user name doesn't match with expected")
		}

		if len(proxy.StackCache) != 1 {
			t.Error("After one successful search in an empty cache, the size of it must be one")
		}

		if proxy.DidLastSearchUsedCache {
			t.Error("No user can be returned from an empty cache")
		}
	})

	t.Run("FindUser - One user, ask for the same user", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}

		if user.ID != knownIDs[0] {
			t.Error("Returned user name doesn't match with expected")
		}

		if len(proxy.StackCache) != 1 {
			t.Error("Cache must not grow if we asked for an object that is stored on it")
		}

		if !proxy.DidLastSearchUsedCache {
			t.Error("The user should have been returned from the cache")
		}
	})

	t.Run("FindUser - Three users", func(t *testing.T) {
		user1, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}

		user2, err := proxy.FindUser(knownIDs[1])
		if err != nil {
			t.Fatal(err)
		}
		if proxy.DidLastSearchUsedCache {
			t.Error("The user wasn't stored on the proxy cache yet")
		}

		user3, err := proxy.FindUser(knownIDs[2])
		if err != nil {
			t.Fatal(err)
		}
		if proxy.DidLastSearchUsedCache {
			t.Error("The user wasn't stored on the proxy cache yet")
		}

		for i := range proxy.StackCache {
			if proxy.StackCache[i].ID == user1.ID {
				t.Errorf("User that should be gone was found")
			}
		}

		if len(proxy.StackCache) != 2 {
			t.Error("After inserting three users, the cache should not grow more than to two")
		}

		for _, v := range proxy.StackCache {
			if v != user2 && v != user3 {
				t.Errorf("A non expected user was found on the cache: %v", v)
			}
		}
	})
}
