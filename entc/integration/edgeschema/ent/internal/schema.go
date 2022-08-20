// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = `{"Schema":"entgo.io/ent/entc/integration/edgeschema/ent/schema","Package":"entgo.io/ent/entc/integration/edgeschema/ent","Schemas":[{"name":"Friendship","config":{"Table":""},"edges":[{"name":"user","type":"User","field":"user_id","unique":true,"required":true,"immutable":true},{"name":"friend","type":"User","field":"friend_id","unique":true,"required":true,"immutable":true}],"fields":[{"name":"weight","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":1,"default_kind":2,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"user_id","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"friend_id","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":3,"MixedIn":false,"MixinIndex":0}}],"indexes":[{"fields":["created_at"]}]},{"name":"Group","config":{"Table":""},"edges":[{"name":"users","type":"User","ref_name":"groups","through":{"N":"joined_users","T":"UserGroup"},"inverse":true}],"fields":[{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":"Unknown","default_kind":24,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}}]},{"name":"Relationship","config":{"Table":""},"edges":[{"name":"user","type":"User","field":"user_id","unique":true,"required":true},{"name":"relative","type":"User","field":"relative_id","unique":true,"required":true},{"name":"info","type":"RelationshipInfo","field":"info_id","unique":true}],"fields":[{"name":"weight","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":1,"default_kind":2,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"user_id","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"relative_id","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"info_id","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":3,"MixedIn":false,"MixinIndex":0}}],"indexes":[{"fields":["weight"]},{"unique":true,"edges":["info"]}],"annotations":{"Fields":{"ID":["user_id","relative_id"],"StructTag":null}}},{"name":"RelationshipInfo","config":{"Table":""},"fields":[{"name":"text","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0}}]},{"name":"Role","config":{"Table":""},"edges":[{"name":"user","type":"User","ref_name":"roles","through":{"N":"roles_users","T":"RoleUser"},"inverse":true}],"fields":[{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"unique":true,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}}]},{"name":"RoleUser","config":{"Table":""},"edges":[{"name":"role","type":"Role","field":"role_id","unique":true,"required":true},{"name":"user","type":"User","field":"user_id","unique":true,"required":true}],"fields":[{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"role_id","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"user_id","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":2,"MixedIn":false,"MixinIndex":0}}],"annotations":{"Fields":{"ID":["user_id","role_id"],"StructTag":null}}},{"name":"Tag","config":{"Table":""},"edges":[{"name":"tweets","type":"Tweet","through":{"N":"tweet_tags","T":"TweetTag"}}],"fields":[{"name":"value","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0}}]},{"name":"Tweet","config":{"Table":""},"edges":[{"name":"liked_users","type":"User","ref_name":"liked_tweets","through":{"N":"likes","T":"TweetLike"},"inverse":true},{"name":"user","type":"User","ref_name":"tweets","through":{"N":"tweet_user","T":"UserTweet"},"inverse":true,"comment":"The uniqueness is enforced on the edge schema"},{"name":"tags","type":"Tag","ref_name":"tweets","through":{"N":"tweet_tags","T":"TweetTag"},"inverse":true}],"fields":[{"name":"text","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":2147483647,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}}]},{"name":"TweetLike","config":{"Table":""},"edges":[{"name":"tweet","type":"Tweet","field":"tweet_id","unique":true,"required":true},{"name":"user","type":"User","field":"user_id","unique":true,"required":true}],"fields":[{"name":"liked_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"user_id","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"tweet_id","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":2,"MixedIn":false,"MixinIndex":0}}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}],"annotations":{"Fields":{"ID":["user_id","tweet_id"],"StructTag":null}}},{"name":"TweetTag","config":{"Table":""},"edges":[{"name":"tag","type":"Tag","field":"tag_id","unique":true,"required":true},{"name":"tweet","type":"Tweet","field":"tweet_id","unique":true,"required":true}],"fields":[{"name":"id","type":{"Type":4,"Ident":"uuid.UUID","PkgPath":"github.com/google/uuid","PkgName":"uuid","Nillable":false,"RType":{"Name":"UUID","Ident":"uuid.UUID","Kind":17,"PkgPath":"github.com/google/uuid","Methods":{"ClockSequence":{"In":[],"Out":[{"Name":"int","Ident":"int","Kind":2,"PkgPath":"","Methods":null}]},"Domain":{"In":[],"Out":[{"Name":"Domain","Ident":"uuid.Domain","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]},"ID":{"In":[],"Out":[{"Name":"uint32","Ident":"uint32","Kind":10,"PkgPath":"","Methods":null}]},"MarshalBinary":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"MarshalText":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"NodeID":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"Time":{"In":[],"Out":[{"Name":"Time","Ident":"uuid.Time","Kind":6,"PkgPath":"github.com/google/uuid","Methods":null}]},"URN":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalBinary":{"In":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"UnmarshalText":{"In":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"Value","Ident":"driver.Value","Kind":20,"PkgPath":"database/sql/driver","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Variant":{"In":[],"Out":[{"Name":"Variant","Ident":"uuid.Variant","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]},"Version":{"In":[],"Out":[{"Name":"Version","Ident":"uuid.Version","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]}}}},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"added_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"tag_id","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"tweet_id","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":3,"MixedIn":false,"MixinIndex":0}}]},{"name":"User","config":{"Table":""},"edges":[{"name":"groups","type":"Group","through":{"N":"joined_groups","T":"UserGroup"}},{"name":"friends","type":"User","through":{"N":"friendships","T":"Friendship"}},{"name":"relatives","type":"User","through":{"N":"relationship","T":"Relationship"}},{"name":"liked_tweets","type":"Tweet","through":{"N":"likes","T":"TweetLike"}},{"name":"tweets","type":"Tweet","through":{"N":"user_tweets","T":"UserTweet"}},{"name":"roles","type":"Role","through":{"N":"roles_users","T":"RoleUser"}}],"fields":[{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":"Unknown","default_kind":24,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}]},{"name":"UserGroup","config":{"Table":""},"edges":[{"name":"user","type":"User","field":"user_id","unique":true,"required":true},{"name":"group","type":"Group","field":"group_id","unique":true,"required":true}],"fields":[{"name":"joined_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"user_id","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"group_id","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":2,"MixedIn":false,"MixinIndex":0}}]},{"name":"UserTweet","config":{"Table":""},"edges":[{"name":"user","type":"User","field":"user_id","unique":true,"required":true},{"name":"tweet","type":"Tweet","field":"tweet_id","unique":true,"required":true}],"fields":[{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"user_id","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"tweet_id","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":2,"MixedIn":false,"MixinIndex":0}}],"indexes":[{"unique":true,"fields":["tweet_id"]}]}],"Features":["entql","sql/upsert","privacy","schema/snapshot"]}`
