module.exports = class Config
  @fetch: (callback) ->
    $.ajax
      type: "GET"
      url: "config.json"
      success: callback