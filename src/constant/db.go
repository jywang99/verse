package constant

const (
    // User
    Id = "id"
    UserTable = "\"user\""
    Name = "name"
    Email = "email"
    Password = "password"
    Created = "created"

    // Collection
    CollectionTable = "collection"
    Path = "path"
    DispName = "disp_name"
    Parent = "parent"

    // Entry
    EntryTable = "entry"
    Desc = "\"desc\""
    ThumbStatic = "thumb_static"
    ThumbDynamic = "thumb_dynamic"
    Updated = "updated"
    Aired = "aired"

    Ascend = "ASC"
    Descend = "DESC"

    // EntryCast
    EntryCastTable = "entry_cast"
    EntryId = "entry_id"
    CastId = "cast_id"
    
    // EntryTag
    EntryTagTable = "entry_tag"
    TagId = "tag_id"

    // Tag
    TagTable = "tag"

    // Cast
    CastTable = "\"cast\""
    PicPath = "pic_path"
)

var UserCols = []string{ Id, Name, Email, Password, Created }
var CollectionCols = []string{ Id, Path, DispName, Parent, Created }
var EntryCols = []string{ Id, Path, DispName, Desc, ThumbStatic, ThumbDynamic, Created, Updated, Aired }
var EntryCastCols = []string{ EntryId, CastId }
var EntryTagCols = []string{ EntryId, TagId }

