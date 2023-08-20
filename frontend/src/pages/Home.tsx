import {
  IonContent,
  IonHeader,
  IonIcon,
  IonItem,
  IonLabel,
  IonList,
  IonMenuButton,
  IonNote,
  IonPage,
  IonTitle,
  IonToolbar,
} from "@ionic/react";
import { chatbubbleOutline, heart } from "ionicons/icons";
import { MockData } from "../store";
import "./Home.css";

const Home = () => {
  const community = MockData[0];

  return (
    <IonPage>
      <IonHeader>
        <IonToolbar>
          <IonTitle>Forum: {community.name}</IonTitle>
          {/* <IonMenuButton i> */}
        </IonToolbar>
      </IonHeader>
      <IonContent fullscreen>
        <IonNote>{community.description}</IonNote>
        <IonList>
          {community.posts.map((post) => (
            <IonItem routerLink={`/${community.name}/${post.id}`} detail>
              <IonLabel>
                <h1>{post.title}</h1>
                <IonNote>{post.body}</IonNote>
              </IonLabel>
              <IonNote slot="end">
                {post.createdAt.toRelative()}
                <br />
                {post.likes}
                <IonIcon icon={heart} />
                {post.commentCount}
                <IonIcon icon={chatbubbleOutline} />
              </IonNote>
            </IonItem>
          ))}
        </IonList>
      </IonContent>
    </IonPage>
  );
};

export default Home;
