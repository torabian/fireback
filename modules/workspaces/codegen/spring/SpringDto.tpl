package com.fireback.modules.{{ .m.Path }};
{{ template "javaimport" . }}
// import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.modules.workspaces.*;
import com.fireback.JsonSerializable;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.Table;

{{ range .children }}
class {{ .FullName }} extends JsonSerializable {
  {{ template "dtoClassContent" . }}
}
{{ end }}


public class {{ .e.DtoName }} extends JsonSerializable {
  {{ template "dtoClassContent" .e }}
}