package repositories

var createPrivateConversationQuery = `
INSERT INTO public."Conversations" ("CreatedBy") VALUES ($1) RETURNING "ID"`

var createGroupConversationQuery = `
INSERT INTO public."Conversations" ("Name", "IsGroup", "CreatedBy") VALUES ($1,$2,$3)`

var addMemberQuery = `
INSERT INTO public."ConversationMembers" ("UserId", "ConversationId") VALUES ($1, $2)`

var getIdPrivateConversationQuery = `
SELECT c."ID"
FROM public."Conversations" c
INNER JOIN public."ConversationMembers" cm1 ON cm1."ConversationId" = c."ID"
INNER JOIN public."ConversationMembers" cm2 ON cm2."ConversationId" = c."ID"
WHERE
  c."IsGroup" = FALSE
  AND cm1."UserId" = $1 
  AND cm2."UserId" = $2
LIMIT 5;
`

var saveChatQuery = `INSERT INTO public."Messages" ("SenderId", "ConversationId", "Content", "MessageType") VALUES ($1,$2,$3,$4)`

var getChatQuery = `
SELECT m."ID", m."Content", m."MessageType", m."CreatedAt", u."ID", u."FullName", u."ProfilePicture" FROM public."Messages" m 
INNER JOIN public."Conversations" c ON c."ID" = m."ConversationId"
INNER JOIN public."Users" u ON u."ID" = m."SenderId"
WHERE c."ID" = $1 
ORDER BY m."CreatedAt" DESC
LIMIT $2 OFFSET $3`

var deleteChatQuery = `DELETE FROM public."Messages" WHERE "ID" = $1`
