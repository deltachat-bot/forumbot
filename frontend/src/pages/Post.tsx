import {
  IonBackButton,
  IonButton,
  IonButtons,
  IonCard,
  IonCardContent,
  IonCardHeader,
  IonCardSubtitle,
  IonCardTitle,
  IonContent,
  IonHeader,
  IonIcon,
  IonList,
  IonModal,
  IonNote,
  IonPage,
  IonText,
  IonTitle,
  IonToolbar,
} from "@ionic/react";
import { RouteComponentProps } from "react-router";
import { MockData, Types } from "../store";
import "./Post.css";
import {
  arrowUndoOutline,
  chatboxOutline,
  chatbubbleOutline,
  heart,
  heartOutline,
} from "ionicons/icons";
import { useRef, useState } from "react";

const MAX_REPLY_DEPTH = 3;

interface PostPageProps
  extends RouteComponentProps<{
    community: string;
    postId: string;
  }> {}

const PostPage: React.FC<PostPageProps> = ({ match }) => {
  const post = MockData.find(
    ({ name }) => name === match.params.community,
  )?.posts.find(({ id }) => id === match.params.postId)!;

  // TODO handle post not found

  return <PostPageInternal post={post} />;
};
export default PostPage;

function PostPageInternal({ post }: { post: Types.Post }) {
  const [replyTargetId, setReplyTargetId] = useState<string | null>(null);

  const onDidDismiss = () => {
    setReplyTargetId(null);
  };
  return (
    <IonPage>
      <IonHeader>
        <IonToolbar>
          <IonButtons slot="start">
            <IonBackButton defaultHref={"/"} />
          </IonButtons>
          <IonTitle>{post.title}</IonTitle>
        </IonToolbar>
      </IonHeader>
      <IonContent fullscreen>
        {/* <IonCard>
          <IonCardHeader>
            <IonCardTitle>{post.title}</IonCardTitle>{" "}
            <IonCardSubtitle>{post.createdAt.toRelative()}</IonCardSubtitle>
          </IonCardHeader>
          <IonCardContent>{post.body}</IonCardContent>
        </IonCard> */}
        <div className="post">
          <h1>{post.title}</h1>
          <IonText>{post.body}</IonText>

          <IonNote style={{ float: "inline-end" }}>
            {post.createdAt.toRelative()}
          </IonNote>
          <ItemActions
            itemId={post.id}
            likeCount={post.likes}
            replyCallback={setReplyTargetId}
          />
        </div>
        <div className="divider">
          <div className="divider-line"></div>
          <IonText color="medium">{post.commentCount} Comments</IonText>
        </div>
        <div className={"post-comments"}>
          {post.comments.map((comment) => (
            <Comment
              comment={comment}
              nestingLevel={0}
              replyCallback={setReplyTargetId}
            />
          ))}
        </div>
        <IonModal
          isOpen={!!replyTargetId}
          initialBreakpoint={0.4}
          breakpoints={[0, 0.8]}
          onDidDismiss={onDidDismiss}
        >
          <div className="reply-modal">
            <h2>
              Reply to{" "}
              {replyTargetId /* todo add some better hint to reply target */}
            </h2>
            <textarea></textarea>
            <br />
            <IonButton>Reply</IonButton>
          </div>
        </IonModal>
      </IonContent>
    </IonPage>
  );
}

function ItemActions({
  itemId,
  likeCount,
  disableReply,
  replyCallback,
}: {
  itemId: string;
  likeCount: number;
  disableReply?: boolean;
  replyCallback: (id: string) => void;
}) {
  const [liked, setLiked] = useState(false);

  // todo load liked

  return (
    <IonButtons>
      <IonButton
        color={liked ? "danger" : "medium"}
        onClick={() => setLiked(!liked)}
      >
        <IonIcon
          icon={liked ? heart : heartOutline}
          aria-label={liked ? "Liked" : "Like"}
        ></IonIcon>
        {likeCount}
      </IonButton>
      {!disableReply && (
        <IonButton
          aria-label={"comment"}
          color="medium"
          onClick={() => replyCallback(itemId)}
        >
          <IonIcon icon={chatboxOutline}></IonIcon> Reply
        </IonButton>
      )}
    </IonButtons>
  );
}

function Comment({
  comment,
  nestingLevel,
  replyCallback,
}: {
  comment: Types.Comment;
  nestingLevel: number;
  replyCallback: (id: string) => void;
}) {
  return (
    <div className={"comment"}>
      <IonText>{comment.body}</IonText>
      <IonNote style={{ float: "inline-end" }}>
        {comment.createdAt.toRelative()}
      </IonNote>
      <ItemActions
        itemId={comment.id}
        likeCount={comment.likes}
        disableReply={nestingLevel >= MAX_REPLY_DEPTH}
        replyCallback={replyCallback}
      />
      <div className={"comments"}>
        {comment.comments.map((comment) => (
          <Comment
            comment={comment}
            nestingLevel={nestingLevel + 1}
            replyCallback={replyCallback}
          />
        ))}
      </div>
    </div>
  );
}
