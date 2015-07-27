package model

// Models can be stored in the database and exported as JSON. Anything that can
// do is a model. Post and Comment.
type Model interface {
   // Sets the LastModified field to the current time. if Timestamp is zero then
   // it sets that too
   Touch()

   // Returns the JSON represention of this model. This should include an "id"
   // field that db.DataSource could use to fetch it from storage. The bytes are
   // encoded in UTF-8.
   ExportJSON() []byte

   // Imports property state from a JSON string. This will wipe the model's data
   // before importing the JSON data. Missing properties will hold their default
   // (zero) value. The []byte should be encoded in UTF-8. Note this will not
   // write any changes to the database.
   ImportJSON([]byte)

   // Update the model by reading a JSON string and retain any property values
   // the JSON ommited. The []byte should be encoded in UTF-8.
   PatchJSON([]byte)
}
