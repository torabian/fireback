package com.fireback.modules.{{ .m.Name }};
{{ template "javaimport" . }}
// import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
{{/* import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel; */}}
import com.fireback.modules.workspaces.*;
import com.fireback.JsonSerializable;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.Table;

{{ range .children }}
@Entity
@Table()
class {{ .FullName }} extends JsonSerializable {
  {{ template "javaClassContent" . }}
}
{{ end }}


@Entity
@Table()
public class {{ .e.EntityName }} extends JsonSerializable {
  {{ template "javaClassContent" .e }}
}