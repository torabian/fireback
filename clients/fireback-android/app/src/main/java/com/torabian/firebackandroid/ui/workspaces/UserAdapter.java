package com.torabian.firebackandroid.ui.workspaces;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.fireback.modules.workspaces.RoleEntity;
import com.fireback.modules.workspaces.UserEntity;
import com.torabian.firebackandroid.R;
import com.torabian.firebackandroid.ui.ListCardItem;

import java.util.List;

public class UserAdapter extends RecyclerView.Adapter<ListCardItem> {
    private List<UserEntity> cardItems;


    public void setCardItems(List<UserEntity> cardItems) {
        this.cardItems = cardItems;
        notifyDataSetChanged();
    }

    @NonNull
    @Override
    public ListCardItem onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
        View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.common_entity_ite, parent, false);
        return new ListCardItem(view);
    }

    @Override
    public void onBindViewHolder(@NonNull ListCardItem holder, int position) {
        UserEntity item = cardItems.get(position);
        holder.bindData(item);
    }

    @Override
    public int getItemCount() {
        return cardItems.size();
    }
}