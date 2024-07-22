// Frabit - The next-generation database automatic operation platform
// Copyright © 2022-2023 Frabit Labs
//
// Licensed under the GNU General Public License, Version 3.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.gnu.org/licenses/gpl-3.0.txt
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/frabits/frabit/pkg/operator"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoType string

const (
	StandaloneMongo MongoType = "STANDALONE"
	ReplicateSet    MongoType = "REPLICATESET"
	SharedCluster   MongoType = "SHAREDCLUSTER"
)

// OperatorMongo is the MongoDB driver.
type OperatorMongo struct {
	DBType  operator.DBType
	connCfg operator.DBConnInfo
	client  *mongo.Client
}

func (op *OperatorMongo) Open(ctx context.Context, dbName operator.DBType, config operator.DBConnInfo) (operator.DBOperator, error) {
	connectionURI := genMongoDBConnectionURI(config)
	opts := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create MongoDB client")
	}
	op.client = client
	op.connCfg = config
	op.DBType = dbName
	return op, nil
}

func (op *OperatorMongo) GetType() operator.DBType {
	return operator.MongoDB
}

func (op *OperatorMongo) Ping(ctx context.Context) error {
	return nil
}

func (op *OperatorMongo) Close(ctx context.Context) error {
	return op.client.Disconnect(ctx)
}

// genMongoDBConnectionURI generate a connection string based provide MongoDB config
func genMongoDBConnectionURI(connCfg operator.DBConnInfo) string {
	connectionURI := "mongodb://"
	if connCfg.SRV {
		connectionURI = "mongodb+srv://"
	}
	if connCfg.User != "" {
		percentEncodingUser := replaceCharacterWithPercentEncoding(connCfg.User)
		percentEncodingPasswd := replaceCharacterWithPercentEncoding(connCfg.Passwd)
		connectionURI = fmt.Sprintf("%s%s:%s@", connectionURI, percentEncodingUser, percentEncodingPasswd)
	}
	connectionURI = fmt.Sprintf("%s%s", connectionURI, connCfg.Host)
	if connCfg.Port != "" {
		connectionURI = fmt.Sprintf("%s:%s", connectionURI, connCfg.Port)
	}
	if connCfg.AuthDB != "" {
		connectionURI = fmt.Sprintf("%s/%s", connectionURI, connCfg.Database)
	}
	// We use admin as the default authentication database.
	// https://www.mongodb.com/docs/manual/reference/connection-string/#mongodb-urioption-urioption.authSource
	authenticationDatabase := connCfg.AuthDB
	if authenticationDatabase == "" {
		authenticationDatabase = "admin"
	}

	if connCfg.Database == "" {
		connectionURI = fmt.Sprintf("%s/", connectionURI)
	}
	connectionURI = fmt.Sprintf("%s?authSource=%s", connectionURI, authenticationDatabase)

	return connectionURI
}

func replaceCharacterWithPercentEncoding(s string) string {
	m := map[string]string{
		":": `%3A`,
		"/": `%2F`,
		"?": `%3F`,
		"#": `%23`,
		"[": `%5B`,
		"]": `%5D`,
		"@": `%40`,
	}
	for k, v := range m {
		s = strings.ReplaceAll(s, k, v)
	}
	return s
}
