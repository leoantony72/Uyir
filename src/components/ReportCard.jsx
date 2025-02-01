import React from "react";
import styles from "./ReportCard.module.css";

export const ReportCard = ({
  type,
  location,
  longitude,
  latitude,
  date,
  status,
  imageUrl,
}) => {
  const statusClassName = `${styles.status} ${
    status === "Fixed" ? styles.statusFixed : styles.statusPending
  }`;

  return (
    <article className={styles.reportCard}>
      <div className={styles.cardContent}>
        <div className={styles.mainSection}>
          <div className={styles.reportDetails}>
            <div className={styles.imageContainer}>
              <img
                className={styles.reportImage}
                src={imageUrl}
                alt={`Road issue: ${type} at ${location}`}
                loading="lazy"
              />
            </div>
            <div className={styles.infoContainer}>
              <div className={styles.reportInfo}>
                <div>Type: {type}</div>
                <div className={styles.location}>Location: {location}</div>
                <div className={styles.coordinates}>
                  <div>Longitude: {longitude}</div>
                  <div>Latitude: {latitude}</div>
                </div>
                <time className={styles.date} dateTime={date}>
                  Date: {date}
                </time>
              </div>
            </div>
          </div>
        </div>
        <div className={styles.statusContainer}>
          <div
            className={statusClassName}
            role="status"
            aria-label={`Report status: ${status}`}
          >
            {status}
          </div>
        </div>
      </div>
    </article>
  );
};
