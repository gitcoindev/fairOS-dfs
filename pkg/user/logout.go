/*
Copyright © 2020 FairOS Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package user

// LogoutUser logs out a giver user from the system and clean him from all the data structures.
func (u *Users) LogoutUser(userName, dataDir, sessionId string) error {
	// basic validations
	if !u.IsUsernameAvailable(userName, dataDir) {
		return ErrInvalidUserName
	}

	// unset cookie and remove user from map
	if !u.IsUserLoggedIn(sessionId) {
		return ErrUserNotLoggedIn
	}

	return u.Logout(sessionId)
}
