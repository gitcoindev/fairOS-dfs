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

import (
	"github.com/fairdatasociety/fairOS-dfs/pkg/blockstore"
	"github.com/fairdatasociety/fairOS-dfs/pkg/feed"
	"github.com/fairdatasociety/fairOS-dfs/pkg/utils"
)

func (*Users) uploadEncryptedMnemonic(userName string, address utils.Address, encryptedMnemonic string, fd *feed.API) error {
	topic := utils.HashString(userName)
	data := []byte(encryptedMnemonic)
	_, err := fd.CreateFeed(topic, address, data)
	return err
}

func (*Users) getEncryptedMnemonic(userName string, address utils.Address, fd *feed.API) (string, error) {
	topic := utils.HashString(userName)
	_, data, err := fd.GetFeedData(topic, address)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (*Users) deleteMnemonic(userName string, address utils.Address, fd *feed.API, client blockstore.Client) error {
	topic := utils.HashString(userName)
	feedAddress, _, err := fd.GetFeedData(topic, address)
	if err != nil {
		return err
	}
	return client.DeleteChunk(feedAddress)
}
