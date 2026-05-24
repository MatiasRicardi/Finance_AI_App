<template>
  <ion-page>
    <ion-tabs>
      <ion-router-outlet />
      <ion-tab-bar slot="bottom">
        <ion-tab-button tab="dashboard" href="/app/dashboard">
          <ion-icon name="home-outline" />
          <ion-label>Dashboard</ion-label>
        </ion-tab-button>
        <ion-tab-button tab="transactions" href="/app/transactions">
          <ion-icon name="list-outline" />
          <ion-label>Transactions</ion-label>
        </ion-tab-button>
        <ion-tab-button tab="import" href="/app/import">
          <ion-icon name="cloud-upload-outline" />
          <ion-label>Import</ion-label>
        </ion-tab-button>
        <ion-tab-button tab="reports" href="/app/reports">
          <ion-icon name="bar-chart-outline" />
          <ion-label>Reports</ion-label>
        </ion-tab-button>
        <ion-tab-button tab="profile" href="/app/profile">
          <ion-icon name="person-outline" />
          <ion-label>Profile</ion-label>
        </ion-tab-button>
      </ion-tab-bar>
    </ion-tabs>
  </ion-page>
</template>

<script setup lang="ts">
import {
  IonPage,
  IonTabs,
  IonRouterOutlet,
  IonTabBar,
  IonTabButton,
  IonIcon,
  IonLabel,
} from '@ionic/vue'
</script>
