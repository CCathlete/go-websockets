-- name: AllPreferences :many
   select id,
          name,
          preference
     from preferences
;

-- name: SetSystemPreference :exec
do $$
begin

delete from preferences where name = $1;

insert into preferences (
    name, preferences, created_at, updated_at
) values ($1, $2, now(), now());

end;
$$
;

-- NOTE: InserOrUpdateSitePreferences is implemented as part of the preferences repo methods, since it's a loop of SetSysetmPreference.
