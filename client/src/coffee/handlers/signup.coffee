module.exports = class Signup
	@save: (form, callback) ->
		$.ajax
			type: "POST"
			url: "signup"
			data: JSON.stringify($(form).serializeObject())
			contentType: "application/json; charset=utf-8"
			dataType: "json"
			success: callback