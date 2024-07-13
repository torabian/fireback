import * as React from "react";
import { DIRECTION, isTreeScrollable } from "../isScrollable";

export interface PullToRefreshProps {
  pullDownContent: JSX.Element;
  releaseContent: JSX.Element;
  refreshContent: JSX.Element;
  pullDownThreshold: number;
  onRefresh: () => Promise<any>;
  triggerHeight?: number | "auto";
  backgroundColor?: string;
  containerStyle?: React.CSSProperties;
  startInvisible?: boolean;
  children?: React.ReactNode;
}

export interface PullToRefreshState {
  pullToRefreshThresholdBreached: boolean;
  maxPullDownDistance: number;
  onRefreshing: boolean;
}

export class PullToRefresh extends React.Component<
  PullToRefreshProps,
  PullToRefreshState
> {
  private container: any;

  private containerRef(container) {
    this.container = container;
  }

  private pullDown: any;

  private pullDownRef(pullDown) {
    this.pullDown = pullDown;
    const maxPullDownDistance =
      this.pullDown &&
      this.pullDown.firstChild &&
      this.pullDown.firstChild["getBoundingClientRect"]
        ? this.pullDown.firstChild["getBoundingClientRect"]().height
        : 0;
    this.setState({ maxPullDownDistance });
  }

  private dragging = false;
  private startY = 0;
  private currentY = 0;

  constructor(props: Readonly<PullToRefreshProps>) {
    super(props);
    this.state = {
      pullToRefreshThresholdBreached: false,
      maxPullDownDistance: 0,
      onRefreshing: false,
    };

    this.containerRef = this.containerRef.bind(this);
    this.pullDownRef = this.pullDownRef.bind(this);
    this.onTouchStart = this.onTouchStart.bind(this);
    this.onTouchMove = this.onTouchMove.bind(this);
    this.onEnd = this.onEnd.bind(this);
  }

  public componentDidMount(): void {
    if (!this.container) {
      return;
    }

    this.container.addEventListener("touchstart", this.onTouchStart);
    this.container.addEventListener("touchmove", this.onTouchMove);
    this.container.addEventListener("touchend", this.onEnd);
    this.container.addEventListener("mousedown", this.onTouchStart);
    this.container.addEventListener("mousemove", this.onTouchMove);
    this.container.addEventListener("mouseup", this.onEnd);
  }

  public componentWillUnmount(): void {
    if (!this.container) {
      return;
    }

    this.container.removeEventListener("touchstart", this.onTouchStart);
    this.container.removeEventListener("touchmove", this.onTouchMove);
    this.container.removeEventListener("touchend", this.onEnd);
    this.container.removeEventListener("mousedown", this.onTouchStart);
    this.container.removeEventListener("mousemove", this.onTouchMove);
    this.container.removeEventListener("mouseup", this.onEnd);
  }

  private onTouchStart(e) {
    const { triggerHeight = 40 } = this.props;
    this.startY = e["pageY"] || e.touches[0].pageY;
    this.currentY = this.startY;

    if (triggerHeight === "auto") {
      const target = e.target;

      const container = this.container;
      if (!container) {
        return;
      }

      // an element we're touching can be scrolled up, so gesture is going to be a scroll gesture
      if (e.type === "touchstart" && isTreeScrollable(target, DIRECTION.up)) {
        return;
      }

      // even though we're not scrolling, the pull-to-refresh isn't visible to the user so cancel
      if (container.getBoundingClientRect().top < 0) {
        return;
      }
    } else {
      const top =
        this.container.getBoundingClientRect().top ||
        this.container.getBoundingClientRect().y ||
        0;
      if (this.startY - top > triggerHeight) {
        return;
      }
    }

    this.dragging = true;
    this.container.style.transition = "transform 0.2s cubic-bezier(0,0,0.31,1)";
    this.pullDown.style.transition = "transform 0.2s cubic-bezier(0,0,0.31,1)";
  }

  private onTouchMove(e) {
    if (!this.dragging) {
      return;
    }

    this.currentY = e["pageY"] || e.touches[0].pageY;
    if (this.currentY < this.startY) {
      return;
    }

    if (e.cancelable) {
      e.preventDefault();
    }

    if (this.currentY - this.startY >= this.props.pullDownThreshold) {
      this.setState({
        pullToRefreshThresholdBreached: true,
      });
    }

    if (this.currentY - this.startY > this.state.maxPullDownDistance) {
      return;
    }

    this.container.style.overflow = "visible";
    this.container.style.transform = `translate(0px, ${
      this.currentY - this.startY
    }px)`;
    this.pullDown.style.visibility = "visible";
  }

  private onEnd() {
    this.dragging = false;
    this.startY = 0;
    this.currentY = 0;

    if (!this.state.pullToRefreshThresholdBreached) {
      this.pullDown.style.visibility = this.props.startInvisible
        ? "hidden"
        : "visible";
      this.initContainer();
      return;
    }

    this.container.style.overflow = "visible";
    this.container.style.transform = `translate(0px, ${this.props.pullDownThreshold}px)`;
    this.setState(
      {
        onRefreshing: true,
      },
      () => {
        this.props.onRefresh().then(() => {
          this.initContainer();
          setTimeout(() => {
            this.setState({
              onRefreshing: false,
              pullToRefreshThresholdBreached: false,
            });
          }, 200);
        });
      },
    );
  }

  private initContainer() {
    requestAnimationFrame(() => {
      if (this.container) {
        this.container.style.overflow = "auto";
        this.container.style.transform = "none";
      }
    });
  }

  private renderPullDownContent() {
    const { releaseContent, pullDownContent, refreshContent, startInvisible } =
      this.props;
    const { onRefreshing, pullToRefreshThresholdBreached } = this.state;
    const content = onRefreshing
      ? refreshContent
      : pullToRefreshThresholdBreached
      ? releaseContent
      : pullDownContent;
    const contentStyle: React.CSSProperties = {
      position: "absolute",
      overflow: "hidden",
      left: 0,
      right: 0,
      top: 0,
      visibility: startInvisible ? "hidden" : "visible",
    };
    return (
      <div id="ptr-pull-down" style={contentStyle} ref={this.pullDownRef}>
        {content}
      </div>
    );
  }

  public render() {
    const { backgroundColor } = this.props;
    const containerStyle: React.CSSProperties = {
      height: "auto",
      overflow: "hidden",
      WebkitOverflowScrolling: "touch",
      position: "relative",
      zIndex: 1,
    };

    if (this.props.containerStyle) {
      Object.keys(this.props.containerStyle).forEach((key: string) => {
        containerStyle[key] = this.props.containerStyle[key];
      });
    }

    if (backgroundColor) {
      containerStyle.backgroundColor = backgroundColor;
    }

    return (
      <div id="ptr-parent" style={containerStyle}>
        {this.renderPullDownContent()}
        <div id="ptr-container" ref={this.containerRef} style={containerStyle}>
          {this.props.children}
        </div>
      </div>
    );
  }
}
