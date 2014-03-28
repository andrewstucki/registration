class Validation
  @phoneMask: ->
    $(".input-mask-phone").mask "(999) 999-9999"

  @chosen: ->
    $(".chzn-select").chosen().change (event) ->
      self = $(event.currentTarget)
      value = self.val() # the select element
      unless value is ""
        $(this).siblings("[class*=\"help-inline\"]:eq(0)").html ""
        $(this).closest(".control-group").removeClass "error"
      else
        $(this).closest(".control-group").addClass "error"
        text = "<span for=\""+self.attr('id')+"\"class=\"help-inline\">
          This field is required.
          </span>"
        $(text).insertAfter $(this).siblings(
          "[class*=\"chzn-container\"]:eq(0)"
        )
    
  @form: (parser, handler) ->
    $("form").validate
      errorElement: "span"
      errorClass: "help-inline"
      focusInvalid: false
      rules: parser.validationRules()
      messages: parser.validationMessages()

      invalidHandler: (event, validator) ->
        $(".alert-error").show()
        false

      highlight: (e) ->
        $(e).closest(".control-group").addClass "error"

      success: (e) ->
        $(e).closest(".control-group").removeClass "error"
        $(e).remove()

      errorPlacement: (error, element) ->
        if element.is(":checkbox") or element.is(":radio")
          controls = element.closest(".controls")
          if controls.find(":checkbox,:radio").length > 1
            controls.append error
          else
            error.insertAfter element.nextAll(".lbl:eq(0)").eq(0)
        else if element.is(".chzn-select")
          error.insertAfter element.siblings(
            "[class*=\"chzn-container\"]:eq(0)"
          )
        else
          error.insertAfter element

      submitHandler: (form) ->
        handler.save form, (data) ->
          $(".alert-error").hide()
          if data.success
            $(".alert-success").show().delay(2000).fadeOut 1000
            $(form)[0].reset()
            $(".chzn-select").val("").trigger "liszt:updated"
          else
            $("#server-error-message").text data.message
            $(".alert-block").show()
        false

class SchemaParser
  constructor: (@schema) ->
    @parseSchema()
  
  parseSchema: ->
    @messages = {}
    @rules = {}
    groups = {}
    @summary_fields = []
    _.each @schema, (options, field) =>
      @summary_fields.push field if options.summarize
      firstKey = if options.group then options.group else field
      groups[firstKey] ?= {}
      groups[firstKey][field] = {}
      _.each options, (option, key) =>
        switch key
          when 'message'
            @messages[field] = option
          when 'required','validation'
            @rules[field] ?= {}
            @rules[field][option] = true if key is 'validation'
            @rules[field].required = option if key is 'required'
          when 'type'
            if not groups[firstKey][field].type
              groups[firstKey][field][key] = option
          when 'formType'
            groups[firstKey][field]["type"] = option
          else
            groups[firstKey][field][key] = option
      groups[firstKey][field].name = field
    @field_groups = _.map(groups, (group, label) ->
      label: label
      fields: group
    )
    console.log @field_groups
    @field_groups
    
  validationMessages: ->
    return @messages if @messages
    @parseSchema()
    @messages
    
  validationRules: ->
    return @rules if @rules
    @parseSchema()
    @rules
    
  fieldGroups: ->
    return @field_groups if @field_groups
    @parseSchema()
    @field_groups
  
  summarizedFields: ->
    return @summary_fields if @summary_fields
    @parseSchema()
    @summary_fields

module.exports.SchemaParser = SchemaParser
module.exports.Validation = Validation