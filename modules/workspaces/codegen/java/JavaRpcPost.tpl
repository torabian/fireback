package com.fireback.modules.{{ .m.Name }};
import com.fireback.ResponseErrorException;
import com.fireback.SingleResponse;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.fireback.ImportRequestDto;
import com.fireback.FirebackConfig;

/*
{{ range $key, $value := .imports }}
{{ if $value.Items}}
  {{ range $value.Items }}
import com.fireback.{{ .}};
  {{ end }}
{{ end }}
{{ end }}
*/

import io.reactivex.rxjava3.core.Single;
import io.reactivex.rxjava3.schedulers.Schedulers;
import okhttp3.MediaType;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.RequestBody;
import okhttp3.Response;
import java.util.concurrent.TimeUnit;

import java.io.IOException;

public class {{ .r.GetFuncNameUpper}} {
    private String getUrl() {
        return FirebackConfig.getInstance().BuildUrl("{{ .r.Url }}");
    }

    public Single<SingleResponse<{{ .r.ResponseEntityComputedSplit }}>> post({{ .r.RequestEntityComputedSplit }} dto) {
        return Single.fromCallable(() -> makeHttpPostRequest(dto))
                .subscribeOn(Schedulers.io());
    }

    private SingleResponse<{{ .r.ResponseEntityComputedSplit }}> makeHttpPostRequest({{ .r.RequestEntityComputedSplit }} dto) throws ResponseErrorException {
        OkHttpClient client = new OkHttpClient.Builder()
            .connectTimeout(10, TimeUnit.SECONDS)
            .writeTimeout(10, TimeUnit.SECONDS)
            .readTimeout(30, TimeUnit.SECONDS)
            .build();

        MediaType mediaType = MediaType.parse("application/json; charset=utf-8");
        RequestBody body = RequestBody.create(mediaType, dto.toJson());
        Request request = new Request.Builder()
                .url(getUrl())
                .post(body)
                .build();

        try (Response response = client.newCall(request).execute()) {
            if (response.isSuccessful()) {
                SingleResponse<{{ .r.ResponseEntityComputedSplit }}> res = SingleResponse.fromJson(response.body().string(), {{ .r.ResponseEntityComputedSplit }}.class);
                response.close();
                return res;
            } else {
                throw ResponseErrorException.fromJson(response.body().string());
            }
        } catch (IOException e) {
            throw ResponseErrorException.fromIoException(e);
        }
    }
}