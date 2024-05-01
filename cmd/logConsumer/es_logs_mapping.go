package main

const esLogsMapping = `{
  "mappings": {
    "properties": {
      "id": {
        "type": "keyword"
      },
      "app": {
        "type": "keyword"
      },
      "system": {
        "type": "keyword"
      },
      "module": {
        "type": "keyword"
      },
      "type": {
        "type": "keyword"
      },
      "pri": {
        "type": "keyword"
      },
      "when": {
        "type": "date"
      },
      "who": {
        "type": "keyword"
      },
      "op": {
        "type": "keyword"
      },
      "class": {
        "type": "keyword"
      },
      "instanceId": {
        "type": "keyword"
      },
      "status": {
        "type": "boolean"
      },
      "error": {
        "type": "text"
      },
      "remote_ip": {
        "type": "ip"
      },
      "msg": {
        "type": "text"
      },
      "data": {
        "properties": {
          "change_data": {
            "properties": {
              "entity": {
                "type": "keyword"
              },
              "op": {
                "type": "keyword"
              },
              "changes": {
                "type": "nested",
                "properties": {
                  "field": {
                    "type": "keyword"
                  },
                  "old_value": {
                    "type": "text"
                  },
                  "new_value": {
                    "type": "text"
                  }
                }
              }
            }
          },
          "activity_data": {
            "type": "text"
          },
          "debug_data": {
            "properties": {
              "pid": {
                "type": "integer"
              },
              "runtime": {
                "type": "keyword"
              },
              "file": {
                "type": "keyword"
              },
              "line": {
                "type": "integer"
              },
              "func": {
                "type": "keyword"
              },
              "stackTrace": {
                "type": "text"
              },
              "data": {
                "type": "object",
                "enabled": false
              }
            }
          }
        }
      }
    }
  }
}`