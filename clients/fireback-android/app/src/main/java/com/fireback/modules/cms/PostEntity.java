package com.fireback.modules.cms;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class PostEntity extends JsonSerializable {
    public String title;
    public String content;
    public PostCategoryEntity category;
    public PostTagEntity[] tags;
    public static class VM extends ViewModel {
    // upper: Title title
    private MutableLiveData< String > title = new MutableLiveData<>();
    public MutableLiveData< String > getTitle() {
        return title;
    }
    public void setTitle( String  v) {
        title.setValue(v);
    }
    // upper: Content content
    private MutableLiveData< String > content = new MutableLiveData<>();
    public MutableLiveData< String > getContent() {
        return content;
    }
    public void setContent( String  v) {
        content.setValue(v);
    }
    // upper: Category category
    private MutableLiveData< PostCategoryEntity > category = new MutableLiveData<>();
    public MutableLiveData< PostCategoryEntity > getCategory() {
        return category;
    }
    public void setCategory( PostCategoryEntity  v) {
        category.setValue(v);
    }
    // upper: Tags tags
    private MutableLiveData< PostTagEntity[] > tags = new MutableLiveData<>();
    public MutableLiveData< PostTagEntity[] > getTags() {
        return tags;
    }
    public void setTags( PostTagEntity[]  v) {
        tags.setValue(v);
    }
    }
}