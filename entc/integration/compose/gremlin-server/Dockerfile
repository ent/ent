# Copyright 2019-present Facebook Inc. All rights reserved.
# This source code is licensed under the Apache 2.0 license found
# in the LICENSE file in the root directory of this source tree.

# Fetch base gremlin server image
# This version compatible with Amazon Neptune Engine Version 1.0.5.0.
# https://docs.aws.amazon.com/neptune/latest/userguide/engine-releases-1.0.5.0.html
FROM tinkerpop/gremlin-server:3.4.11

# Copy overridden server configuration.
COPY gremlin-server.yaml tinkergraph-empty.properties /opt/gremlin-server/conf/
