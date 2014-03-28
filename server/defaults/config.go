package defaults

var Config = `{
  "title": "NSWN Registration!",
  "success": "You are all done now, enjoy the program!",
  "invalid": "Please correct the issues below!",
  "error": "Please notify one of the staff."
}`

var Schema = `{
  "first": {
    "type": "text",
    "required": true,
    "group": "name",
    "label": "First",
    "summarize": true
  },
  "last": {
    "type": "text",
    "required": true,
    "group": "name",
    "label": "Last",
    "summarize": true
  },
  "gender": {
    "type": "enum",
    "values": ["Male", "Female"],
    "required": true,
    "formType": "radio",
    "message": "Please choose a gender.",
    "summarize": true
  },
  "dorm": {
    "type": "enum",
    "values": [
      "Centennial",
      "Territorial",
      "Frontier",
      "Pioneer",
      "Yudof",
      "Comstock",
      "Middlebrook",
      "Sanford",
      "Bailey",
      "Wilkins",
      "University Village",
      "Grand Marc",
      "University Commons",
      "Stadium View",
      "17th Ave.",
      "Other"
    ],
    "required": true,
    "formType": "select",
    "message": "Please choose your location.",
    "summarize": true
  },
  "year": {
    "type": "enum",
    "values": [
      "Freshman",
      "Sophomore",
      "Junior",
      "Junior Transfer",
      "Senior",
      "Graduate"
    ],
    "required": true,
    "formType": "radio",
    "message": "Please choose your year in school",
    "summarize": true
  },
  "email": {
    "type": "text",
    "required": true,
    "message": "Please provide a valid email.",
    "validation": "email",
    "prepend": "icon-envelope"
  },
  "phone": {
    "type": "text",
    "required": false,
    "validation": "phone",
    "prepend": "icon-phone",
    "extraClass": "input-mask-phone"
  },
  "church": {
    "type": "boolean",
    "group": "interests",
    "label": "Looking for a Church"
  },
  "smallgroup": {
    "type": "boolean",
    "group": "interests",
    "label": "Small Groups"
  },
  "bible": {
    "type": "boolean",
    "group": "interests",
    "label": "Large Group Bible Studies"
  }
}`
