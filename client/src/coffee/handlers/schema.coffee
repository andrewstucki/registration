module.exports = class Schema
  @fetch: (callback) ->
    $.ajax
      type: "GET"
      url: "schema.json"
      success: callback