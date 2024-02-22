package com.torabian.firebackandroid.ui;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.fireback.modules.workspaces.RoleEntity;
import com.torabian.firebackandroid.R;


import java.util.List;

public class CustomAdapter extends RecyclerView.Adapter<ListCardItem> {
    private List<RoleEntity> cardItems;


    public void setCardItems(List<RoleEntity> cardItems) {
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
        RoleEntity item = cardItems.get(position);
        holder.bindData(item);
    }

    @Override
    public int getItemCount() {
        return cardItems.size();
    }
}