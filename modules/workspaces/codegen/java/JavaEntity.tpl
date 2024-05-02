package com.fireback.modules.{{ .m.Path }};
{{ template "javaimport" . }}
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;

{{ range .children }}
class {{ .FullName }} extends JsonSerializable {
  {{ template "definitionrow" .CompleteFields }}


  public static class VM extends ViewModel {
    {{ template "viewmodelrow" .CompleteFields }}
  }
}
{{ end }}


public class {{ .e.EntityName }} extends JsonSerializable {
    {{ template "definitionrow" .e.CompleteFields }}

    public static class VM extends ViewModel {
      {{ template "viewmodelrow" .e.CompleteFields }}
    }
}