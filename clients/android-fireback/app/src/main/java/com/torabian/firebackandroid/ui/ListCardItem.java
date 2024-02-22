package com.torabian.firebackandroid.ui;

import android.view.View;
import android.widget.TextView;

import com.fireback.modules.workspaces.RoleEntity;
import com.torabian.firebackandroid.R;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

public class ListCardItem extends RecyclerView.ViewHolder {

    private TextView titleTextView;
    private TextView dateTextView;

    public ListCardItem(@NonNull View itemView) {
        super(itemView);
        titleTextView = itemView.findViewById(R.id.card_item_h1);
        dateTextView = itemView.findViewById(R.id.card_item_date);
        // Initialize other UI elements for the card as needed
    }

    public void bindData(RoleEntity role) {
        titleTextView.setText(role.name);
        dateTextView.setText(role.createdFormatted);
    }
}
